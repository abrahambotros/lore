# LORE
Light Object-Relational Environment (LORE) provides a very simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go

[![GoDoc](https://godoc.org/github.com/abrahambotros/lore?status.svg)](https://godoc.org/github.com/abrahambotros/lore) [![Build Status](https://travis-ci.org/abrahambotros/lore.svg)](https://travis-ci.org/abrahambotros/lore)


## Motivation
With LORE, you weave your own lore and control your own magic (... bear with me, I'll explain). LORE provides a thin veil that abstracts away some of the inconveniences of object-relational mapping in general, but intentionally and explicitly avoids doing any hand-wavy magic tricks that place someone else's mystery black box between you and your data.

To do so, LORE acts as a simple layer gluing together some wonderful and powerful libraries, providing a minimal abstraction above them for convenience:

* Squirrel (https://github.com/Masterminds/squirrel) for SQL generation.
  * In particular, the entrypoints for building queries via LORE (Query.BuildSqlSelect/Insert/Update/Delete and related Query methods) directly return squirrel SQL builders, so you get all of the power of Squirrel with some added convenience (tablename handling, common related queries, etc.) via LORE.
* SQLX (http://jmoiron.github.io/sqlx/) for running DB SQL queries and parsing results.

Aside from this, you're in charge of your own schema, migration, etc. - no full ORM and cruft to get in the way of your lore!

Lastly, while I'm sure it could still be improved, I've tried (and will continue to try) to ensure documentation provides full transparency of what is going on, so that your lore is entirely in your control.

## Quick-start and examples
Below is a quick run-through of getting set up and constructing some simple examples. Check GoDoc documentation for more information on any item, and look through the test source files for more thorough examples of usage. I also recommend using string constants everywhere for your db field names (string literals everywhere are literally the devil and should be banned from your lore IMO - pun intended), but am just using literals here ONLY for simplicity and succinctness.

### Config
When initializing your app, you should inform LORE of your desired config before constructing any LORE queries. To do so, you can use the `SetConfig` function and pass in a `*lore.Config`. Currently, this only determines the SQL placeholder format used in all Squirrel queries.

```go
lore.SetConfig(&lore.Config{
    SQLPlaceholderFormat: lore.SQLPlaceholderFormatDollar, // Or: SQLPlaceholderFormatQuestion
})

// Or you can use the default config, which currently sets to the dollar-sign/$# placeholder format,
// but it's probably best to call SetConfig explicitly as above.
lore.SetConfigDefault()
```

### Model interface
To use LORE with your model, your model should implement `lore.ModelInterface`; see an example of
this below. Note that you can use "db:" tags for sqlx directly; LORE currently doesn't do any
specific handling of tags, as it delegates this all to sqlx.

Please check out the GoDoc for [ModelInterface](https://godoc.org/github.com/abrahambotros/lore#ModelInterface) for much more detail for each method you need to implement, in particular `DbFieldMap` (which shouldn't include auto-managed keys, such as serial keys).

```go
type Legend struct {
    Id      int     `db:"id"` // Serial/auto-incrementing primary key
    Name    string  `db:"name"`
    Culture string  `db:"culture"`
}

var _ lore.ModelInterface = (*Legend)(nil)

func (*Legend) DbTableName() string {
    return "legends"
}

func (l *Legend) DbFieldMap() map[string]interface{} {
    // Note that Id is purposefully left out, since this is a serial/auto-incrementing field!
    return map[string]interface{}{
        "name": l.Name,
        "culture": l.Culture,
    }
}

func (*Legend) DbPrimaryFieldKey() string {
    return "id"
}

func (l *Legend) DbPrimaryFieldValue() interface{} {
    return l.Id
}
```

### Build SQL query
Now we can use LORE to build a SQL query for your model. This involves 3 steps:

1. Create new `*lore.Query` instance, passing in your model (implementing `ModelInterface`).
2. Build SQL query via the `lore.Query.BuildSql*` methods, which wrap Squirrel builders.
3. Pass the completed SQL builder you made in step (2) back to the Query from (1) via `lore.Query.SetSqlBuilder`. Note that this can be combined with step (2) in simple cases as in the example below.

Let's try building the following query on the `Legend` model we defined above:

`SELECT * FROM legends WHERE name = 'Mars' AND culture = 'Roman' LIMIT 1;`

```go
/*
1. Create new *lore.Query with our Legend model.
*/
q := lore.NewQuery(&Legend{})

/*
2. Build the SQL query; here we use the BuildSqlSelectStar entrypoint and finish building via
Squirrel, then set it back into the Query via SetSqlBuilder in the same command. This tells LORE
that this is the SQl that will be run when the Query is executed later.

Note that BuildSqlSelectStar is an example of just one convenience wrapper that LORE provides;
alternatively, you can just build the SQL via Squirrel, or use BuildSqlSelect and pass in your
own column names, etc.
*/
q.SetSqlBuilder(
    q.BuildSqlSelectStar(). // This returns a Squirrel builder directly now, so the rest of this chain here is purely Squirrel.
    Where(squirrel.Eq{
        "name": "Mars",
        "culture": "Roman",
    }).
    Limit(1),
)

// If you want, you can use ToSql to get the Squirrel ToSql representation of the Query at any time.
qSql, qArgs, err := q.ToSql()
```

### Execute SQL query on your DB
Now that we've built a `*lore.Query` and attached our SQL builder to it, we can execute this as a query against our DB. LORE currently provides 3 methods for doing so:

1. `*lore.Query.Execute` - wraps sqlx.DB.Exec
2. `*lore.Query.ExecuteThenParseSingle` - wraps sqlx.DB.Get
3. `*lore.Query.ExecuteThenParseList` - wraps sqlx.DB.List

All of these methods attempt to return the number of rows affected by the query, along with of course an error if one was encountered. See GoDoc for more details, especially regarding numRowsAffected (see comments for `Execute` in particular).

```go
/*
Execute and parse to list. Note that we pass in a POINTER to a list of structs we want to
scan the DB rows back into - when passing into the Execute* Query methods, LORE assumes you're
passing in a pointer (either to a list or a single struct), and will return an error if it
detects otherwise.
*/
db := getDb() // ... Your own lore should conjure up a *sqlx.DB instance here.
discoveredLegend := &Legend{}
// See notes for numRowsAffected in Query.Execute documentation.
numRowsAffected, err := q.ExecuteThenParseSingle(db, discoveredLegend)
if err != nil {
    // Handle errors here.
}
// Row matching the SQL query is written into discoveredLegend. Do whatever with it now.
discoveredLegend.GetCorrespondingCelestialBody()
discoveredLegend.TellUsYourTale()

...

/*
Also try the other query wrappers listed below for creating your SQL... (See GoDoc for more
details and possibly more functions)
*/
q.SetSqlBuilder(
    q.BuildSqlSelect(columns ...string)
    q.BuildSqlSelectStar()

    q.BuildSqlInsert()
    q.BuildSqlInsertColumnsAndValues(columns []string, values []interface{})
    q.BuildSqlInsertModel()

    q.BuildSqlUpdate()
    q.BuildSqlUpdateModelByPrimaryKey()

    q.BuildSqlDelete()
    q.BuildSqlDeleteModelByPrimaryKey()
)
```

## DANGER: WIP
LORE is a major WIP. Contributions are welcome, but use in production is cautioned against at the moment unless you know full well what you're doing!

## TODO
* Allow using sqlx QueryRow/QueryRowX for large/unrestricted-length queries instead of just Get/Select.
* Better tests, especially for Execute\* methods.
* Consider better way to relate updates to SQL-builders to the parent query without having to call `SetSqlBuilder` every time.
* Dedicated examples in GoDoc.

## Final notes
Thanks for looking, and please feel free to message me if you're having any problems with LORE! I'm also always open to suggestions on ways to improve LORE, whether minor changes/fixes or even large rewrites - always happy to learn of better ways to do things!
