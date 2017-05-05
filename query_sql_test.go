package lore

import (
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/Masterminds/squirrel"
)

func TestSqlSelectBuilder(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlSelect(STAR).
			Where(squirrel.Eq{
				"f1": 1,
			}).
			Where(squirrel.Eq{
				"f2": 2,
			}).
			Limit(10).
			Suffix(RETURNING_STAR),
	)

	// Test intermediate SQL.
	sql, sqlArgs, err := q.ToSql()
	if err != nil {
		t.Fatalf("Error building query to SQL: %s", err.Error())
	}
	if sql != fmt.Sprintf("SELECT * FROM %s WHERE f1 = $1 AND f2 = $2 LIMIT 10 RETURNING *", _TEST_DB_TABLENAME) {
		t.Fatalf("Incorrect SQL output: %s", sql)
	}
	if len(sqlArgs) != 2 || sqlArgs[0] != 1 || sqlArgs[1] != 2 {
		t.Fatalf("Incorrect SQL args: %+v", sqlArgs)
		return
	}

	// Test execute.
	// TODO: Improve testing of query results.
	db, dbMock := getTestSqlxDb(t)
	dbMock.ExpectQuery(fmt.Sprintf("^SELECT \\* FROM %s.*", _TEST_DB_TABLENAME)).WithArgs(1, 2).WillReturnRows(sqlmock.NewRows([]string{}))
	numRowsAffected, err := q.ExecuteThenParseList(db, newTestModelEmptyList())
	if err != nil {
		t.Fatalf("Error in ExecuteThenParseList: %s", err.Error())
	}
	if numRowsAffected != 0 {
		t.Fatalf("Unexpected numRowsAffected (%d) != 0", numRowsAffected)
	}
}
