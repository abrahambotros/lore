package lore

// import "github.com/Masterminds/squirrel"

// /*
// SqlInsertBuilder specializes a SqlBuilder for INSERT queries.
// */
// type SqlInsertBuilder struct {
// 	SqlBuilder               SqlBuilder
// 	SquirrelStatementBuilder squirrel.InsertBuilder
// }

// /*
// Enforce SqlBuilderInterface interface.
// */
// var _ SqlBuilderInterface = (*SqlInsertBuilder)(nil)

// /*
// Insert provides the entrypoint for specializing a generic Query as an INSERT query.
// */
// func (q *Query) Insert(columns []string, values []interface{}) *SqlInsertBuilder {
// 	return &SqlInsertBuilder{
// 		SqlBuilder: SqlBuilder{
// 			query: q,
// 		},
// 		SquirrelStatementBuilder: newSquirrelStatementBuilder().
// 			Insert(q.modelInterface.DbTableName()).
// 			Columns(columns...).
// 			Values(values...),
// 	}
// }

// /*
// ToSql builds the corresponding SQL and args for this query.
// */
// func (sib *SqlInsertBuilder) ToSql() (sql string, args []interface{}, err error) {
// 	return sib.SquirrelStatementBuilder.ToSql()
// }
