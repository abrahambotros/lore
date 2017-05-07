package lore

import (
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

/*
SelectModelByPrimaryKey builds and executes a SQL statement on the db similar to the following SQL,
scanning the result into resultSinglePtr. This is essentially a convenience wrapper around
BuildSqlSelectStar and ExecuteThenParseSingle.

`SELECT * FROM <table> WHERE <primary field key> = <primary field value> LIMIT 1;`
*/
func SelectModelByPrimaryKey(mi ModelInterface, db *sqlx.DB, resultSinglePtr interface{}) (found bool, err error) {
	// Ensure non-empty primary field key.
	if mi.DbPrimaryFieldKey() == "" {
		return false, errors.New(_ERR_EMPTY_PRIMARY_KEY)
	}

	// Build query.
	q := NewQuery(mi)
	q.SetSqlBuilder(
		q.BuildSqlSelectStar().
			Where(squirrel.Eq{
				mi.DbPrimaryFieldKey(): mi.DbPrimaryFieldValue(),
			}).
			Limit(1),
	)

	// Execute ParseSingle query to return single, if found.
	return q.ExecuteThenParseSingle(db, resultSinglePtr)
}

/*
SelectModelWhere builds and executes a SQL statement on the db similar to the following SQL,
scanning the result into resultSinglePtr. This is essentially a convenience wrapper around
BuildSqlSelectStar and ExecuteThenParseSingle, applying the WHERE clause accordingly.

Note that where is a pointer; if no where is desired, pass nil instead. When desired, this is
typically something like a squirrel.Eq instance, etc.

`SELECT * FROM <table> WHERE <where conditions> LIMIT 1;`
*/
func SelectModelWhere(mi ModelInterface, db *sqlx.DB, where *sqlPart, resultSinglePtr interface{}) (found bool, err error) {
	// Build query.
	q := NewQuery(mi)
	qSqlBuilder := q.BuildSqlSelectStar()
	// Add where if non-nil.
	if where != nil {
		qSqlBuilder = qSqlBuilder.Where(*where)
	}
	q.SetSqlBuilder(qSqlBuilder)

	// Execute ParseSingle query.
	return q.ExecuteThenParseSingle(db, resultSinglePtr)
}

/*
SelectModelsWhere is analogous to SelectModelWhere, but wraps ExecuteThenParseList instead of
ExecuteThenParseSingle, and allows appying a LIMIT clause too.

Note that limit is a pointer here - if nil is supplied, no limit is set on the SQL statement;
otherwise, the underlying limit uint64 is applied.

`SELECT * FROM <table> WHERE <where conditions> LIMIT <limit>;`
*/
func SelectModelsWhere(mi ModelInterface, db *sqlx.DB, where *sqlPart, limit *uint64, resultListPtr interface{}) (numRowsAffected uint64, err error) {
	// Build query.
	q := NewQuery(mi)
	qSqlBuilder := q.BuildSqlSelectStar()
	// Add where if non-nil.
	if where != nil {
		qSqlBuilder = qSqlBuilder.Where(*where)
	}
	// Add limit if non-nil.
	if limit != nil {
		qSqlBuilder = qSqlBuilder.Limit(*limit)
	}
	q.SetSqlBuilder(qSqlBuilder)

	// Execute ParseList query.
	return q.ExecuteThenParseList(db, resultListPtr)
}

/*
InsertNewModel builds and executes a SQL statement on the db similar to the following SQL, scanning
the result into resultSinglePtr. This is essentially a convenience wrapper around
BuildSqlInsertModel and ExecuteThenParseSingle.

`INSERT INTO <table> (<columns from DbFieldMap>) VALUES (<values from DbFieldMap>) RETURNING * ;`
*/
func InsertNewModel(mi ModelInterface, db *sqlx.DB, resultSinglePtr interface{}) error {
	// Build query.
	q := NewQuery(mi)
	qSqlBuilder, err := q.BuildSqlInsertModel()
	if err != nil {
		return err
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Execute ParseSingle query.
	_, err = q.ExecuteThenParseSingle(db, resultSinglePtr)
	return err
}

/*
UpdateModelByPrimaryKey builds and executes a SQL statement on the db similar to the following SQL,
scanning the result into resultSinglePtr. This is essentially a convenience wrapper around
BuildSqlUpdateModelByPrimaryKey and ExecuteThenParseSingle.

`UPDATE <table> SET <columns and values from DbFieldMap>
	WHERE <primary field key> = <primary field value> RETURNING * ;`
*/
func UpdateModelByPrimaryKey(mi ModelInterface, db *sqlx.DB, resultSinglePtr interface{}) (found bool, err error) {
	// Build query.
	q := NewQuery(mi)
	qSqlBuilder, err := q.BuildSqlUpdateModelByPrimaryKey()
	if err != nil {
		return false, err
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Execute ParseSingle query.
	return q.ExecuteThenParseSingle(db, resultSinglePtr)
}

/*
UpdateSetMapWhere builds and executes a SQL statement on the db similar to the following SQL,
scanning the result into resultListPtr. This is essentially a convenience wrapper around
BuildSqlUpdateSetMap and ExecuteThenParseList, applying the WHERE clause accordingly.

Note that where is a pointer; if no where is desired, pass nil instead. When desired, this is
typically something like a squirrel.Eq instance, etc.

`UPDATE <table> SET <columns and values from map> WHERE <where conditions> RETURNING * ;`
*/
func UpdateSetMapWhere(mi ModelInterface, db *sqlx.DB, m map[string]interface{}, where *sqlPart, resultListPtr interface{}) (numRowsAffected uint64, err error) {
	// Build query.
	q := NewQuery(mi)
	qSqlBuilder := q.BuildSqlUpdateSetMap(m)
	// Add where if non-nil.
	if where != nil {
		qSqlBuilder = qSqlBuilder.Where(*where)
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Execute ParseList query.
	return q.ExecuteThenParseList(db, resultListPtr)
}

/*
DeleteModelByPrimaryKey builds and executes a SQL statement on the db similar to the following SQL,
scanning the result into resultSinglePtr. This is essentially a convenience wrapper around
BuildSqlDeleteModelByPrimaryKey and ExecuteThenParseSingle.

`DELETE FROM <table> WHERE <primary field key> = <primary field value> RETURNING * ;`
*/
func DeleteModelByPrimaryKey(mi ModelInterface, db *sqlx.DB, resultSinglePtr interface{}) (found bool, err error) {
	// Build query.
	q := NewQuery(mi)
	qSqlBuilder, err := q.BuildSqlDeleteModelByPrimaryKey()
	if err != nil {
		return false, err
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Execute ParseSingle query.
	return q.ExecuteThenParseSingle(db, resultSinglePtr)
}

/*
DeleteModelsWhere builds and executes a SQL statement on the db similar to the following SQL,
scanning the result into resultListPtr. This is essentially a convenience wrapper around
BuildSqlDelete and ExecuteThenParseList, applying the WHERE clause accordingly.

Note that where is a pointer; if no where is desired, pass nil instead. When desired, this is
typically something like a squirrel.Eq instance, etc.

`DELETE FROM <table> WHERE <where conditions> RETURNING * ;`
*/
func DeleteModelsWhere(mi ModelInterface, db *sqlx.DB, where *sqlPart, resultListPtr interface{}) (numRowsAffected uint64, err error) {
	// Build query.
	q := NewQuery(mi)
	qSqlBuilder := q.BuildSqlDelete()
	// Add where if non-nil.
	if where != nil {
		qSqlBuilder = qSqlBuilder.Where(*where)
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Execute ParseList query.
	return q.ExecuteThenParseList(db, resultListPtr)
}
