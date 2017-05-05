package lore

// import (
// 	"github.com/Masterminds/squirrel"
// )

// /*
// SqlSelectBuilder specializes a SqlBuilder for SELECT queries.
// */
// type SqlSelectBuilder struct {
// 	// SqlBuilder               SqlBuilder
// 	SquirrelStatementBuilder squirrel.SelectBuilder
// }

// /*
// Enforce SqlBuilderInterface interface.
// */
// var _ SqlBuilderInterface = (*SqlSelectBuilder)(nil)

// /*
// Select provides the entrypoint for specializing a generic Query as a QuerySelectBuilder.
// */
// func (q *Query) Select(columns []string) *Query {
// 	// Build SqlSelectBuilder.
// 	sqlSelectBuilder := &SqlSelectBuilder{
// 		// SqlBuilder: SqlBuilder{
// 		// 	query: q,
// 		// },
// 		SquirrelStatementBuilder: newSquirrelStatementBuilder().
// 			Select(columns...).
// 			From(q.modelInterface.DbTableName()),
// 	}

// 	// Attach to parent query.
// 	q.sqlBuilder = sqlSelectBuilder

// 	// Return parent query.
// 	return q
// }

// /*
// ToSql builds the corresponding SQL and args for this query.
// */
// func (ssb *SqlSelectBuilder) ToSql() (sql string, args []interface{}, err error) {
// 	return ssb.SquirrelStatementBuilder.ToSql()
// }

// /*
// Where ...
// */
// func (ssb *SqlSelectBuilder) Where(pred interface{}, args ...interface{}) SqlBuilderInterface {
// 	ssb.SquirrelStatementBuilder = ssb.SquirrelStatementBuilder.Where(pred, args...)
// 	return ssb
// }

// /*
// Limit ...
// */
// func (ssb *SqlSelectBuilder) Limit(limit int) SqlBuilderInterface {
// 	ssb.SquirrelStatementBuilder = ssb.SquirrelStatementBuilder.Limit(limit)

// }
