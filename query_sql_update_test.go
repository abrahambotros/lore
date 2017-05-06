package lore

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
)

func TestBuildSqlUpdate(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueId := 1
	tmTestValueField := tm.Field*2 + 1

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlUpdate().
			Set(_TEST_DB_FIELDNAME_FIELD, tmTestValueField).
			Where(squirrel.Eq{
				_TEST_DB_FIELDNAME_ID: tmTestValueId,
			}).
			Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	err := testBuildSqlHelper(
		q,
		fmt.Sprintf(
			"UPDATE %s SET %s = $1 WHERE %s = $2 RETURNING *",
			_TEST_DB_TABLENAME,
			_TEST_DB_FIELDNAME_FIELD,
			_TEST_DB_FIELDNAME_ID,
		),
		[]interface{}{tmTestValueField, tmTestValueId},
	)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBuildSqlUpdateModelByPrimaryKey(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueId := 1
	tmTestValueField := tm.Field*2 + 1
	tm.Id = tmTestValueId
	tm.Field = tmTestValueField

	// Build query and SQL.
	q := NewQuery(tm)
	qSqlBuilder, err := q.BuildSqlUpdateModelByPrimaryKey()
	if err != nil {
		t.Error(err)
		return
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	err = testBuildSqlHelper(
		q,
		fmt.Sprintf(
			"UPDATE %s SET %s = $1 WHERE %s = $2 RETURNING *",
			_TEST_DB_TABLENAME,
			_TEST_DB_FIELDNAME_FIELD,
			_TEST_DB_FIELDNAME_ID,
		),
		[]interface{}{tm.Field, tm.Id},
	)
	if err != nil {
		t.Error(err)
		return
	}
}
