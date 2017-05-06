package lore

import (
	"fmt"
	"testing"
)

func TestBuildSqlInsertColumnsAndValues(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueField := tm.Field*2 + 1

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlInsertColumnsAndValues(
			[]string{_TEST_DB_FIELDNAME_FIELD},
			[]interface{}{tmTestValueField},
		).Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES ($1) RETURNING *",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELD,
	)
	expectedArgs := []interface{}{tmTestValueField}
	err := testBuildSqlHelper(q, expectedSql, expectedArgs)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBuildSqlInsertModel(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueField := tm.Field*2 + 1
	tm.Field = tmTestValueField

	// Build query and SQL.
	q := NewQuery(tm)
	qSqlBuilder, err := q.BuildSqlInsertModel()
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
			"INSERT INTO tests (%s) VALUES ($1) RETURNING *",
			_TEST_DB_FIELDNAME_FIELD,
		),
		[]interface{}{tm.Field},
	)
	if err != nil {
		t.Error(err)
		return
	}
}
