package lore

// import "github.com/Masterminds/squirrel"

// /*
// SqlUpdateBuilder specializes a SqlBuilder for UPDATE queries.
// */
// type SqlUpdateBuilder struct {
// 	SqlBuilder               SqlBuilder
// 	SquirrelStatementBuilder squirrel.UpdateBuilder
// }

// /*
// Enforce SqlBuilderInterface interface.
// */
// var _ SqlBuilderInterface = (*SqlUpdateBuilder)(nil)

// /*
// Update provides the entrypoint for specializing a generic Query as an UPDATE query.
// */
// func (q *Query) Update(setMap map[string]interface{}) *SqlUpdateBuilder {
// 	return &SqlUpdateBuilder{
// 		SqlBuilder: SqlBuilder{
// 			query: q,
// 		},
// 		SquirrelStatementBuilder: newSquirrelStatementBuilder().
// 			Update(q.modelInterface.DbTableName()).
// 			SetMap(setMap),
// 	}
// }

// /*
// ToSql builds the corresponding SQL and args for this query.
// */
// func (sub *SqlUpdateBuilder) ToSql() (sql string, args []interface{}, err error) {
// 	return sub.SquirrelStatementBuilder.ToSql()
// }
