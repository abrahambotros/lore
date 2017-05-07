package lore

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

/*
part is directly imported from Squirrel.
*/
type part struct {
	pred interface{}
	args []interface{}
}

/*
newPart is directly imported from Squirrel.
*/
func newPart(pred interface{}, args ...interface{}) squirrel.Sqlizer {
	return &part{pred, args}
}

/*
ToSql is directly imported from Squirrel.
*/
func (p part) ToSql() (sql string, args []interface{}, err error) {
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
