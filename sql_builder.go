package lore

// /*
// SqlBuilder provides a base type for SQL builders.
// */
// type SqlBuilder struct {
// 	query *Query
// }

/*
SqlBuilderInterface provides a generic interface for SQL-izable structs. This mirrors
squirrel.Sqlizer.
*/
type SqlBuilderInterface interface {
	ToSql() (sql string, args []interface{}, err error)
	// Where(pred interface{}, args ...interface{}) SqlBuilderInterface
	// Limit(limit int) SqlBuilderInterface
}
