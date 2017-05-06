package lore

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
)

func TestBuildSqlSelect(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	testValueWhereFieldOne := 1
	testValueWhereFieldTwo := "Two"
	limit := uint64(10)

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlSelect(_TEST_DB_FIELDNAME_FIELDONE, _TEST_DB_FIELDNAME_FIELDTWO).
			Where(squirrel.Eq{
				_TEST_DB_FIELDNAME_FIELDONE: testValueWhereFieldOne,
			}).
			Where(squirrel.Eq{
				_TEST_DB_FIELDNAME_FIELDTWO: testValueWhereFieldTwo,
			}).
			Limit(limit),
	)

	// Delegate to SQL test helper.
	expectedSql := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s = $1 AND %s = $2 LIMIT %d",
		fmt.Sprintf("%s, %s", _TEST_DB_FIELDNAME_FIELDONE, _TEST_DB_FIELDNAME_FIELDTWO),
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELDONE,
		_TEST_DB_FIELDNAME_FIELDTWO,
		limit,
	)
	expectedArgs := []interface{}{testValueWhereFieldOne, testValueWhereFieldTwo}
	testBuildSqlHelper(t, q, expectedSql, expectedArgs)
}
