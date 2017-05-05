package lore

// import (
// 	"github.com/Masterminds/squirrel"
// )

// /*
// SqlDeleteBuilder specializes a SqlBuilder for DELETE queries.
// */
// type SqlDeleteBuilder struct {
// 	SqlBuilder               SqlBuilder
// 	SquirrelStatementBuilder squirrel.DeleteBuilder
// }

// /*
// Enforce SqlBuilderInterface interface.
// */
// var _ SqlBuilderInterface = (*SqlDeleteBuilder)(nil)

// /*
// Delete provides the entrypoint for specializing a generic Query as a DELETE query.
// */
// func (q *Query) Delete() *SqlDeleteBuilder {
// 	return &SqlDeleteBuilder{
// 		SqlBuilder: SqlBuilder{
// 			query: q,
// 		},
// 		SquirrelStatementBuilder: newSquirrelStatementBuilder().
// 			Delete(q.modelInterface.DbTableName()),
// 	}
// }

// /*
// ToSql builds the corresponding SQL and args for this query.
// */
// func (sdb *SqlDeleteBuilder) ToSql() (sql string, args []interface{}, err error) {
// 	return sdb.SquirrelStatementBuilder.ToSql()
// }
