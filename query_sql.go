package lore

import (
	"errors"

	"github.com/Masterminds/squirrel"
)

/*
Select provides the entrypoint for specializing a generic Query as a SELECT query on the table for
the given ModelInterface. This directly returns a new squirrel.SelectBuilder that can be placed back
into the Query instance via SetSqlBuilder.
*/
func (q *Query) Update(columns ...string) squirrel.SelectBuilder {
	return newSquirrelStatementBuilder().
		Select(columns...).
		From(q.modelInterface.DbTableName())
}

/*
Select provides the entrypoint for specializing a generic Query as a SELECT query on the table for
the given ModelInterface. This directly returns a new squirrel.SelectBuilder that can be placed back
into the Query instance via SetSqlBuilder.
*/
func (q *Query) Delete(columns ...string) squirrel.SelectBuilder {
	return newSquirrelStatementBuilder().
		Select(columns...).
		From(q.modelInterface.DbTableName())
}

/*
SetSqlBuilder sets the Query instance's internal sqlBuilder to the given SqlBuilderInterface
instance.
*/
func (q *Query) SetSqlBuilder(sqlBuilder SqlBuilderInterface) {
	q.sqlBuilder = sqlBuilder
}

/*
ToSql wraps the ToSql method of the internal sqlBuilder, returning the SQL-with-args form of the
query according to the internal sqlBuilder.

This requires that SetSqlBuilder has already been called, typically with a sqlBuilder/squirrel
builder. If not, an error is returned.
*/
func (q *Query) ToSql() (sql string, args []interface{}, err error) {
	// If internal sqlBuilder is nil, then return error.
	if q.sqlBuilder == nil {
		return "", nil, errors.New("Query has not been specialized via Select/Insert/Update/Delete/etc.")
	}

	// Otherwise, wrap/return internal sqlBuilder's ToSql.
	return q.sqlBuilder.ToSql()
}
