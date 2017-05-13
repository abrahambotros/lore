package lore

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

/*
Where returns a new part as a Sqlizer interface, allowing this to be called in typical WHERE
clauses via Squirrel.
*/
func Where(pred interface{}, args ...interface{}) *sqlPart {
	return newSqlPart(pred, args...)
}

/*
BuildWhereColumnEqualsSubquerySelect returns a new Where-compatible *sqlPart as a Sqlizer interface,
allowing this to be called in typical WHERE clauses via Squirrel. When the returned *sqlPart is
supplied to a Squirrel.Where call, this produces SQL similar to:
`WHERE <column> = (<subquery>)`

TODO: Tests.

TODO: Update docs above. Note doesn't work if already have prefix.
*/
func BuildWhereColumnEqualsSubquerySelect(column string, subquerySqlBuilder squirrel.SelectBuilder) squirrel.SelectBuilder {
	return subquerySqlBuilder.
		Prefix(
			fmt.Sprintf("WHERE %s = (", column),
		).
		Suffix(")")

	// // Get SQL from subquery.
	// subSql, subArgs, err := subquery.ToSql()
	// if err != nil {
	// 	return nil, err
	// }

	// // Wrap subSql with equality to column.
	// subSql = fmt.Sprintf("%s = (%s)", column, subSql)

	// // Return/wrap as Where/*sqlPart instance.
	// return Where(subSql, subArgs), nil

	// Add prefix.
}
