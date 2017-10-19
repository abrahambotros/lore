package lore

import (
	"fmt"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/jmoiron/sqlx"
)

/*
getTestSqlxDb returns a sqlx.DB for testing, along with a mock handle.
*/
func getTestSqlxDb(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error getting test mock sqlx db: %s", err.Error())
	}
	sqlxDb := sqlx.NewDb(mockDb, "sqlmock")
	return sqlxDb, mock
}

/*
TestExecuteUnexported tests the (unexported) Query execute method.
*/
func TestExecuteUnexported(t *testing.T) {
	db, _ := getTestSqlxDb(t)

	// Build test query.
	q := NewQuery(newTestModelEmpty())
	q.SetSqlBuilder(
		q.BuildSqlSelectStar(),
	)

	// Test nil db resilience.
	numRowsAffected, err := q.execute(nil, nil, _EXECUTE_MODE_NO_PARSE)
	if numRowsAffected != 0 || err == nil {
		t.Errorf("Expected 0 rows affected and non-empty err since nil db")
		return
	}

	// Test invalid mode resilience.
	numRowsAffected, err = q.execute(db, nil, 42)
	if numRowsAffected != 0 || err == nil {
		t.Errorf("Expected 0 rows affected and non-empty err since invalid mode")
		return
	}
}

/*
TestExecute tests the Query Execute method using a mock db.
*/
func TestExecute(t *testing.T) {
	// Build test query.
	q := NewQuery(newTestModelEmpty())
	q.SetSqlBuilder(
		q.BuildSqlSelectStar(),
	)

	// Build test/mock db.
	db, dbMock := getTestSqlxDb(t)
	dbMock.ExpectExec(fmt.Sprintf("^SELECT \\* FROM %s", _TEST_DB_TABLENAME)).WillReturnResult(sqlmock.NewResult(1, 1))
	numRowsAffected, err := q.Execute(db)
	if err != nil {
		t.Errorf("Error in Execute: %s", err.Error())
		return
	}
	if numRowsAffected != 1 {
		t.Errorf("Unexpected numRowsAffected (%d) != 1", numRowsAffected)
		return
	}
	err = dbMock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
		return
	}
}

/*
TestExecuteThenParseSingle tests the Query ExecuteThenParseSingle method using a mock db.

TODO: Better testing of return result.
TODO: Test invalid model interface.
*/
func TestExecuteThenParseSingle(t *testing.T) {
	// Build test query.
	q := NewQuery(newTestModelEmpty())
	q.SetSqlBuilder(
		q.BuildSqlSelectStar().Limit(1).Suffix(RETURNING_STAR),
	)

	// Build test/mock db.
	db, dbMock := getTestSqlxDb(t)
	dbMock.ExpectQuery(fmt.Sprintf("^SELECT \\* FROM %s LIMIT 1 RETURNING \\*", _TEST_DB_TABLENAME)).
		WillReturnRows(sqlmock.NewRows([]string{}))
	found, err := q.ExecuteThenParseSingle(db, newTestModelEmpty())
	if err != nil {
		t.Errorf("Error in ExecuteThenParseList: %s", err.Error())
		return
	}
	if found {
		t.Errorf("Unexpected found = %t", found)
		return
	}

	// Should get error if pass in empty list to ExecuteThenParseSingle.
	_, err = q.ExecuteThenParseSingle(db, newTestModelEmptyList())
	if err == nil {
		t.Error("Expected error from passing in list instead of single to ExecuteThenParseSingle, but got no such error")
		return
	}
}

/*
TestExecuteThenParseList tests the Query ExecuteThenParseList method using a mock db.

TODO: Better testing of return result.
*/
func TestExecuteThenParseList(t *testing.T) {
	// Build test query.
	q := NewQuery(newTestModelEmpty())
	q.SetSqlBuilder(
		q.BuildSqlSelectStar().Suffix(RETURNING_STAR),
	)

	// Build test/mock db.
	db, dbMock := getTestSqlxDb(t)
	dbMock.ExpectQuery(fmt.Sprintf("^SELECT \\* FROM %s RETURNING \\*", _TEST_DB_TABLENAME)).
		WillReturnRows(sqlmock.NewRows([]string{}))
	numRowsAffected, err := q.ExecuteThenParseList(db, newTestModelEmptyList())
	if err != nil {
		t.Errorf("Error in ExecuteThenParseList: %s", err.Error())
		return
	}
	if numRowsAffected != 0 {
		t.Errorf("Unexpected numRowsAffected (%d) != 0", numRowsAffected)
		return
	}

	// Should get error if pass in empty-but-not-list to ExecuteThenParseList.
	_, err = q.ExecuteThenParseList(db, newTestModelEmpty())
	if err == nil {
		t.Error("Expected error from passing in single instead of list to ExecuteThenParseList, but got no such error")
		return
	}
}
