package lore

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

/*
getTestSqlxDb ...
*/
func getTestSqlxDb(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error getting test mock sqlx db: %s", err.Error())
	}
	sqlxDb := sqlx.NewDb(mockDb, "sqlmock")
	return sqlxDb, mock
}
