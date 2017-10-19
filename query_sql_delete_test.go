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
	err := testBuildSqlHelper(q, expectedSql, expectedArgs)
	if err != nil {
		t.Error(err)
		return
	}
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
		t.Error(err)
		return
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
	err = testBuildSqlHelper(q, expectedSql, expectedArgs)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBuildSqlDeleteModelByPrimaryKeyInvalid(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInvalidInstance()
	q := NewQuery(tm)

	// Ensure that primary key is empty for this invalid instance.
	if tm.DbPrimaryFieldKey() != "" {
		t.Error("Expected empty primary field key for testing")
		return
	}

	// Expect error since empty primary key.
	_, err := q.BuildSqlDeleteModelByPrimaryKey()
	if err == nil {
		t.Error("Expected non-nil err since invalid empty primary key", err)
		return
	}
}
