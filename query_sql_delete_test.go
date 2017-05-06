package lore

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
)

func TestBuildSqlDelete(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueId := 1

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlDelete().
			Where(squirrel.Eq{
				_TEST_DB_FIELDNAME_ID: tmTestValueId,
			}).
			Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"DELETE FROM %s WHERE %s = $1 RETURNING *",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_ID,
	)
	expectedArgs := []interface{}{tmTestValueId}
	testBuildSqlHelper(t, q, expectedSql, expectedArgs)
}

func TestBuildSqlDeleteModelByPrimaryKey(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueId := 1
	tm.Id = tmTestValueId

	// Build query and SQL.
	q := NewQuery(tm)
	qSqlBuilder, err := q.BuildSqlDeleteModelByPrimaryKey()
	if err != nil {
		t.Fatal(err)
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"DELETE FROM %s WHERE %s = $1 RETURNING *",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_ID,
	)
	expectedArgs := []interface{}{tmTestValueId}
	testBuildSqlHelper(t, q, expectedSql, expectedArgs)

}
