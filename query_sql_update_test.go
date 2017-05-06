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
	tmTestValueFieldOne := tm.FieldOne*2 + 1
	tmTestValueFieldTwo := fmt.Sprintf("%s-%s", tm.FieldTwo, "Two")

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlUpdate().
			Set(_TEST_DB_FIELDNAME_FIELDONE, tmTestValueFieldOne).
			Set(_TEST_DB_FIELDNAME_FIELDTWO, tmTestValueFieldTwo).
			Where(squirrel.Eq{
				_TEST_DB_FIELDNAME_ID: tmTestValueId,
			}).
			Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper, but manually handle args permutations since the last arg is
	// fixed.
	var expectedSql string
	var expectedArgs []interface{}
	expectedSqlBase := "UPDATE %s SET %s = $1, %s = $2 WHERE %s = $3 RETURNING *"
	var err error
	for i := 0; i < 2; i++ {
		if i == 0 {
			// Forward order for first 2 args, 3rd fixed.
			expectedSql = fmt.Sprintf(
				expectedSqlBase,
				_TEST_DB_TABLENAME,
				_TEST_DB_FIELDNAME_FIELDONE,
				_TEST_DB_FIELDNAME_FIELDTWO,
				_TEST_DB_FIELDNAME_ID,
			)
			expectedArgs = []interface{}{tmTestValueFieldOne, tmTestValueFieldTwo, tmTestValueId}
		} else {
			// Reverse order for first 2 args, 3rd fixed.
			expectedSql = fmt.Sprintf(
				expectedSqlBase,
				_TEST_DB_TABLENAME,
				_TEST_DB_FIELDNAME_FIELDTWO,
				_TEST_DB_FIELDNAME_FIELDONE,
				_TEST_DB_FIELDNAME_ID,
			)
			expectedArgs = []interface{}{tmTestValueFieldTwo, tmTestValueFieldOne, tmTestValueId}
		}
		err = testBuildSqlHelper(q, expectedSql, expectedArgs)
		if err == nil {
			// If no error, then break immediately.
			break
		}
	}
	// If still err here, then all permutations failed.
	if err != nil {
		t.Errorf("All SQL permutations failed. Last expectedSql: %s. Last expectedArgs: %+v", expectedSql, expectedArgs)
		return
	}
}

func TestBuildSqlUpdateModelByPrimaryKey(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	tmTestValueId := 1
	tmTestValueFieldOne := tm.FieldOne*2 + 1
	tmTestValueFieldTwo := fmt.Sprintf("%s-%s", tm.FieldTwo, "Two")
	tm.Id = tmTestValueId
	tm.FieldOne = tmTestValueFieldOne
	tm.FieldTwo = tmTestValueFieldTwo

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

	// Delegate to SQL test helper, but manually handle args permutations since the last arg is
	// fixed.
	var expectedSql string
	var expectedArgs []interface{}
	expectedSqlBase := "UPDATE %s SET %s = $1, %s = $2 WHERE %s = $3 RETURNING *"
	for i := 0; i < 2; i++ {
		if i == 0 {
			// Forward order for first 2 args, 3rd fixed.
			expectedSql = fmt.Sprintf(
				expectedSqlBase,
				_TEST_DB_TABLENAME,
				_TEST_DB_FIELDNAME_FIELDONE,
				_TEST_DB_FIELDNAME_FIELDTWO,
				_TEST_DB_FIELDNAME_ID,
			)
			expectedArgs = []interface{}{tm.FieldOne, tm.FieldTwo, tm.Id}
		} else {
			// Reverse order for first 2 args, 3rd fixed.
			expectedSql = fmt.Sprintf(
				expectedSqlBase,
				_TEST_DB_TABLENAME,
				_TEST_DB_FIELDNAME_FIELDTWO,
				_TEST_DB_FIELDNAME_FIELDONE,
				_TEST_DB_FIELDNAME_ID,
			)
			expectedArgs = []interface{}{tm.FieldTwo, tm.FieldOne, tm.Id}
		}
		err = testBuildSqlHelper(q, expectedSql, expectedArgs)
		if err == nil {
			// If no error, then break immediately.
			break
		}
	}
	// If still err here, then all permutations failed.
	if err != nil {
		t.Errorf("All SQL permutations failed. Last expectedSql: %s. Last expectedArgs: %+v", expectedSql, expectedArgs)
		return
	}
}
