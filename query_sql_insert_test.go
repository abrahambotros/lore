package lore

import (
	"fmt"
	"testing"
)

func TestBuildSqlInsertColumnsAndValues(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueFieldOne := tm.FieldOne*2 + 1
	tmTestValueFieldTwo := fmt.Sprintf("%s-%s", tm.FieldTwo, "Two")

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlInsertColumnsAndValues(
			[]string{_TEST_DB_FIELDNAME_FIELDONE, _TEST_DB_FIELDNAME_FIELDTWO},
			[]interface{}{tmTestValueFieldOne, tmTestValueFieldTwo},
		).Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"INSERT INTO %s (%s,%s) VALUES ($1,$2) RETURNING *",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELDONE,
		_TEST_DB_FIELDNAME_FIELDTWO,
	)
	expectedArgs := []interface{}{tmTestValueFieldOne, tmTestValueFieldTwo}
	testBuildSqlHelper(t, q, expectedSql, expectedArgs)
}

func TestBuildSqlInsertModel(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueFieldOne := tm.FieldOne*2 + 1
	tmTestValueFieldTwo := fmt.Sprintf("%s-%s", tm.FieldTwo, "Two")
	tm.FieldOne = tmTestValueFieldOne
	tm.FieldTwo = tmTestValueFieldTwo

	// Build query and SQL.
	q := NewQuery(tm)
	qSqlBuilder, err := q.BuildSqlInsertModel()
	if err != nil {
		t.Fatal(err)
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"INSERT INTO %s (%s,%s) VALUES ($1,$2) RETURNING *",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELDONE,
		_TEST_DB_FIELDNAME_FIELDTWO,
	)
	expectedArgs := []interface{}{tm.FieldOne, tm.FieldTwo}
	testBuildSqlHelper(t, q, expectedSql, expectedArgs)
}
