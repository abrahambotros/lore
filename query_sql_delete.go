package lore

import (
	"errors"

	"github.com/Masterminds/squirrel"
)

/*
BuildSqlDelete provides the entrypoint for specializing a generic Query as a DELETE query on the
table for the given ModelInterface. This directly returns a new squirrel.DeleteBuilder that can be
placed back into the Query instance via SetSqlBuilder; the underlying SQL has the form:
"DELETE FROM <DbTableName>".
*/
func (q *Query) BuildSqlDelete() squirrel.DeleteBuilder {
	return newSquirrelStatementBuilder().
		Delete(q.modelInterface.DbTableName())
}

/*
BuildSqlDeleteModelByPrimaryKey wraps BuildSqlDelete to perform the delete on the table row with the
matching primary key/value of this Query's ModelInterface's model instance. Alias for
query.BuildSqlDelete().Where(<primary key and value>).
*/
func (q *Query) BuildSqlDeleteModelByPrimaryKey() (squirrel.DeleteBuilder, error) {
	mi := q.modelInterface

	// Return error if primary key is empty.
	pk := mi.DbPrimaryFieldKey()
	if pk == "" {
		return q.BuildSqlDelete(), errors.New("Empty primary key encountered; cannot perform delete by primary key")
	}

	return q.BuildSqlDelete().
		Where(squirrel.Eq{
			pk: mi.DbPrimaryFieldValue(),
		}), nil
}
