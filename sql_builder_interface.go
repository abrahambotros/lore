package lore

/*
SqlBuilderInterface provides a generic interface for SQL-izing. This mirrors the squirrel.Sqlizer
interface.
*/
type SqlBuilderInterface interface {
	ToSql() (sql string, args []interface{}, err error)
}
