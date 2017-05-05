package lore

import "github.com/Masterminds/squirrel"

/*
SquirrelStatementBuilder provides a generic interface for squirrel statement builders.
*/
type SquirrelStatementBuilder interface {
	ToSql() (sql string, args []interface{}, err error)
}

/*
newSquirrelStatementBuilder returns a new squirrel.StatementBuilder(Type) instance, with the current
config applied.
*/
func newSquirrelStatementBuilder() squirrel.StatementBuilderType {
	c := GetConfig()
	return squirrel.StatementBuilder.
		PlaceholderFormat(c.SQLPlaceholderFormat)
}
