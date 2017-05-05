package lore

/*
Query provides a generic, chainable query instance for callers to configure a specific query.
*/
type Query struct {
	modelInterface ModelInterface
	sqlBuilder     SqlBuilderInterface
	// db             *sqlx.DB
}

/*
NewQuery instantiates a new Query instance based on the given ModelInterface.
*/
func NewQuery(modelInterface ModelInterface) *Query {
	return &Query{
		modelInterface: modelInterface,
		sqlBuilder:     nil,
		// db:             db,
	}
}

// /*
// ToSql wraps the ToSql method of the internal sqlBuilder, returning the SQL-with-args form of the
// query according to the internal sqlBuilder.

// This requires that methods like Select/Insert/Update/Delete have already been called on the query,
// specializing the query. If one of these methods has not been called, an error is returned here.
// */
// func (q *Query) ToSql() (sql string, args []interface{}, err error) {
// 	// If internal sqlBuilder is nil, then return error.
// 	if q.sqlBuilder == nil {
// 		return "", nil, errors.New("Query has not been specialized via Select/Insert/Update/Delete/etc.")
// 	}

// 	// Otherwise, wrap/return internal sqlBuilder's ToSql.
// 	return q.sqlBuilder.ToSql()
// }

// /*
// Where ...

// TODO: DOCUMENT
// */
// func (q *Query) Where(pred interface{}, args ...interface{}) *Query {
// 	q.sqlBuilder = q.sqlBuilder.Where(pred, args...)
// 	return q
// }

// /*
// Limit ...

// TODO: DOCUMENT
// */
// func (q *Query) Limit(limit int) *Query {
// 	q.sqlBuilder = q.sqlBuilder.Limit(limit)
// }
