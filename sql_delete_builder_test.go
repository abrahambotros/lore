package lore

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/jmoiron/sqlx"
// )

// func TestSqlDeleteBuilder(t *testing.T) {
// 	SetConfigDefault()
// 	tm := createTestModelInstance()

// 	sql, args, err := NewQuery(tm, &sqlx.DB{}).
// 		Delete().
// 		Where(LtOrEq{
// 			"f1": 1,
// 		}).
// 		ToSql()

// 	if err != nil {
// 		t.Fatalf("Error building query to SQL: %s", err.Error())
// 	}
// 	if sql != fmt.Sprintf(
// 		"DELETE FROM %s WHERE f1 <= $1",
// 		_TEST_DB_TABLENAME,
// 	) {
// 		t.Fatalf("Incorrect SQL output: %s", sql)
// 	}
// 	if len(args) != 1 || args[0] != 1 {
// 		t.Fatalf("Incorrect args: %+v", args)
// 	}
// }
