package lore

import "github.com/Masterminds/squirrel"

/*
BuildSqlInsert provides the entrypoint for specializing a generic Query as an INSERT query on the
table for the given ModelInterface. This directly returns a new squirrel.InsertBuilder that can be
placed back into the Query instance via SetSqlBuilder; the underlying SQL has the form:
"INSERT INTO <DbTableName>".
*/
func (q *Query) BuildSqlInsert() squirrel.InsertBuilder {
	return newSquirrelStatementBuilder().
		Insert(q.modelInterface.DbTableName())
}

/*
BuildSqlInsertColumnsAndValues wraps BuildSqlInsert with the given columns and values. Alias for
query.BuildSqlInsert().Columns(<columns...>).Values(<values...>).
*/
func (q *Query) BuildSqlInsertColumnsAndValues(columns []string, values []interface{}) squirrel.InsertBuilder {
	return q.BuildSqlInsert().Columns(columns...).Values(values...)
}

/*
BuildSqlInsertModel wraps BuildSqlInsert with the given ModelInterface's DbFieldMap as the INSERT
query's columns and values. Uses the ModelInterface the Query was originally created with. Alias for
query.BuildSqlInsertColumnsAndValues(<model's DbFieldMap keys>, <model's DbFieldMap values>).
*/
func (q *Query) BuildSqlInsertModel() (squirrel.InsertBuilder, error) {
	// Parse columns and values from model's DbFieldMap.
	columns, values, err := getMapKeysVals(q.modelInterface.DbFieldMap())
	if err != nil {
		return q.BuildSqlInsert(), err
	}

	// Delegate to BuildSqlInsertColumnsAndValues.
	return q.BuildSqlInsertColumnsAndValues(columns, values), nil
}
