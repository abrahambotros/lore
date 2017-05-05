package lore

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/jmoiron/sqlx"
// )

// func TestSqlUpdateBuilder(t *testing.T) {
// 	SetConfigDefault()
// 	tm := createTestModelInstance()
// 	newFieldTwo := "Three"
// 	newFieldOne := 2
// 	whereFieldOne := 10

// 	sql, args, err := NewQuery(tm, &sqlx.DB{}).
// 		Update(
// 			map[string]interface{}{
// 				_TEST_DB_FIELDNAME_FIELDONE: newFieldOne,
// 				_TEST_DB_FIELDNAME_FIELDTWO: newFieldTwo,
// 			},
// 		).
// 		Where(Eq{
// 			_TEST_DB_FIELDNAME_FIELDONE: whereFieldOne,
// 		}).
// 		ToSql()

// 	if err != nil {
// 		t.Fatalf("Error building query to SQL: %s", err.Error())
// 	}
// 	if sql != fmt.Sprintf(
// 		"UPDATE %s SET %s = $1, %s = $2 WHERE %s = $3",
// 		_TEST_DB_TABLENAME, _TEST_DB_FIELDNAME_FIELDONE, _TEST_DB_FIELDNAME_FIELDTWO, _TEST_DB_FIELDNAME_FIELDONE,
// 	) {
// 		t.Fatalf("Incorrect SQL output: %s", sql)
// 	}
// 	if len(args) != 3 || args[0] != newFieldOne || args[1] != newFieldTwo || args[2] != whereFieldOne {
// 		t.Fatalf("Incorrect args: %+v", args)
// 	}
// }
