package lore

// import "github.com/Masterminds/squirrel"

// /*
// Type wrappers around squirrel types. Note that we do not use pointers for exact aliases as this
// would prevent single-command use of created instances when passing to Where; for example, we would
// have to do:

// 	eq := Eq{}
// 	...Where(*eq)

// instead of simply:

// 	...Where(Eq{})

// */
// type Eq squirrel.Eq
// type NotEq squirrel.NotEq
// type Lt squirrel.Lt
// type LtOrEq squirrel.LtOrEq
// type Gt squirrel.Gt
// type GtOrEq squirrel.GtOrEq
// type Or squirrel.Or
// type GenericSqlizer struct {
// 	sql  string
// 	args []interface{}
// }

// /*
// Provide wrapper ToSql method for all expressions above, allowing these to be treated as
// squirrel.Sqlizers and inherit all relevant functionality.
// */
// func (eq Eq) ToSql() (sql string, args []interface{}, err error) {
// 	return squirrel.Eq(eq).ToSql()
// }
// func (notEq NotEq) ToSql() (sql string, args []interface{}, err error) {
// 	return squirrel.NotEq(notEq).ToSql()
// }
// func (lt Lt) ToSql() (sql string, args []interface{}, err error) {
// 	return squirrel.Lt(lt).ToSql()
// }
// func (ltOrEq LtOrEq) ToSql() (sql string, args []interface{}, err error) {
// 	return squirrel.LtOrEq(ltOrEq).ToSql()
// }
// func (gt Gt) ToSql() (sql string, args []interface{}, err error) {
// 	return squirrel.Gt(gt).ToSql()
// }
// func (gtOrEq GtOrEq) ToSql() (sql string, args []interface{}, err error) {
// 	return squirrel.GtOrEq(gtOrEq).ToSql()
// }
// func (or Or) ToSql() (sql string, args []interface{}, err error) {
// 	return squirrel.Or(or).ToSql()
// }
// func (gs GenericSqlizer) ToSql() (sql string, args []interface{}, err error) {
// 	return gs.sql, gs.args, nil
// }
