package lore

import "github.com/Masterminds/squirrel"

/*
BuildSqlSelect provides the entrypoint for specializing a generic Query as a SELECT query on the table
for the given ModelInterface. This directly returns a new squirrel.SelectBuilder that can be placed
back into the Query instance via SetSqlBuilder; the underlying SQL has the form:
"SELECT <columns> FROM <DbTableName>".
*/
func (q *Query) BuildSqlSelect(columns ...string) squirrel.SelectBuilder {
	return newSquirrelStatementBuilder().
		Select(columns...).
		From(q.modelInterface.DbTableName())
}

/*
BuildSqlSelectStar wraps BuildSqlSelect as a SELECT * query. Alias for query.BuildSqlSelect("*").
*/
func (q *Query) BuilSqlSelectStar() squirrel.SelectBuilder {
	return q.BuildSqlSelect(STAR)
}
