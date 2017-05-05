package lore

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/jmoiron/sqlx"
// )

// func TestSqlInsertBuilder(t *testing.T) {
// 	tm := createTestModelInstance()
// 	newFieldTwo := "Three"
// 	newFieldOne := 2
// 	SetConfigDefault()

// 	sql, args, err := NewQuery(tm, &sqlx.DB{}).
// 		Insert(
// 			[]string{_TEST_DB_FIELDNAME_FIELDTWO, _TEST_DB_FIELDNAME_FIELDONE},
// 			[]interface{}{newFieldTwo, newFieldOne},
// 		).ToSql()

// 	if err != nil {
// 		t.Fatalf("Error building query to SQL: %s", err.Error())
// 	}
// 	if sql != fmt.Sprintf(
// 		"INSERT INTO %s (%s,%s) VALUES ($1,$2)",
// 		_TEST_DB_TABLENAME, _TEST_DB_FIELDNAME_FIELDTWO, _TEST_DB_FIELDNAME_FIELDONE,
// 	) {
// 		t.Fatalf("Incorrect SQL output: %s", sql)
// 	}
// 	if len(args) != 2 || args[0] != newFieldTwo || args[1] != newFieldOne {
// 		t.Fatalf("Incorrect args: %+v", args)
// 	}
// }
