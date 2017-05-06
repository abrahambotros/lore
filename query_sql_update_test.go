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

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"UPDATE %s SET %s = $1, %s = $2 WHERE %s = $3 RETURNING *",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELDONE,
		_TEST_DB_FIELDNAME_FIELDTWO,
		_TEST_DB_FIELDNAME_ID,
	)
	expectedArgs := []interface{}{tmTestValueFieldOne, tmTestValueFieldTwo, tmTestValueId}
	testBuildSqlHelper(t, q, expectedSql, expectedArgs)
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
		t.Fatal(err)
	}
	q.SetSqlBuilder(
		qSqlBuilder.Suffix(RETURNING_STAR),
	)

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"UPDATE %s SET %s = $1, %s = $2 WHERE %s = $3 RETURNING *",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELDONE,
		_TEST_DB_FIELDNAME_FIELDTWO,
		_TEST_DB_FIELDNAME_ID,
	)
	expectedArgs := []interface{}{tmTestValueFieldOne, tmTestValueFieldTwo, tmTestValueId}
	testBuildSqlHelper(t, q, expectedSql, expectedArgs)

}
