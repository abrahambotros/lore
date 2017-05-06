package lore

import (
	"errors"

	"github.com/Masterminds/squirrel"
)

/*
BuildSqlUpdate provides the entrypoint for specializing a generic Query as an UPDATE query on the
table for the given ModelInterface. This directly returns a new squirrel.UpdateBuilder that can be
placed back into the Query instance via SetSqlBuilder; the underlying SQL has the form:
"UPDATE <DbTableName>".
*/
func (q *Query) BuildSqlUpdate() squirrel.UpdateBuilder {
	return newSquirrelStatementBuilder().
		Update(q.modelInterface.DbTableName())
}

/*
BuildSqlUpdateModelByPrimaryKey wraps BuildSqlUpdate to perform the update with the given columns
and values defined by the Query's ModelInterface's DbFieldMap on the table row with the matching
primary key/value for this ModelInterface's model instance. Alias for
query.BuildSqlUpdate().Where(<primary key and value>).Set(<columns and values according to DbFieldMap>).
*/
func (q *Query) BuildSqlUpdateModelByPrimaryKey() (squirrel.UpdateBuilder, error) {
	mi := q.modelInterface

	// Return error if primary key is empty.
	pk := mi.DbPrimaryFieldKey()
	if pk == "" {
		return q.BuildSqlUpdate(), errors.New("Empty primary key encountered; cannot perform update by primary key")
	}

	return q.BuildSqlUpdate().
		Where(squirrel.Eq{
			pk: mi.DbPrimaryFieldValue(),
		}).
		SetMap(mi.DbFieldMap()), nil
}
