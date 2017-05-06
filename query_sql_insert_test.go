package lore

import (
	"fmt"
	"testing"
)

func TestBuildSqlInsertColumnsAndValues(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueFieldTwo := fmt.Sprintf("%s-%s", tm.FieldTwo, "Two")

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlInsertColumnsAndValues(
			[]string{_TEST_DB_FIELDNAME_FIELDTWO},
			[]interface{}{tmTestValueFieldTwo},
		).Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES ($1) RETURNING *",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELDTWO,
	)
	expectedArgs := []interface{}{tmTestValueFieldTwo}
	err := testBuildSqlHelper(q, expectedSql, expectedArgs)
	if err != nil {
		t.Error(err)
		return
	}
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
		t.Error(err)
		return
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper. Note that we need to test the different permutations of the
	// arguments.
	expectedSqlBase := "INSERT INTO tests (%s,%s) VALUES ($1,$2) RETURNING *"
	testBuildSqlHelperWithArgPermutations(
		t,
		q,
		expectedSqlBase,
		[]string{_TEST_DB_FIELDNAME_FIELDONE, _TEST_DB_FIELDNAME_FIELDTWO},
		[]interface{}{tm.FieldOne, tm.FieldTwo},
	)
}
