package lore

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

/*
sqlPart is directly imported from Squirrel.part.
*/
type sqlPart struct {
	pred interface{}
	args []interface{}
}

/*
newSqlPart is directly imported from Squirrel.newPart.
*/
func newSqlPart(pred interface{}, args ...interface{}) *sqlPart {
	return &sqlPart{pred, args}
}

/*
ToSql is directly imported from Squirrel.part.ToSql.
*/
func (p sqlPart) ToSql() (sql string, args []interface{}, err error) {
	switch pred := p.pred.(type) {
	case nil:
		// no-op
	case squirrel.Sqlizer:
		sql, args, err = pred.ToSql()
	case string:
		sql = pred
		args = p.args
	default:
		err = fmt.Errorf("expected string or Sqlizer, not %T", pred)
	}
	return
}
