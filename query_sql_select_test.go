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
				_TEST_DB_FIELDNAME_FIELDTWO: testValueWhereFieldTwo,
			}).
			Limit(limit),
	)

	// Delegate to SQL test helper with args permutations.
	testBuildSqlHelperWithArgPermutations(
		t,
		q,
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE %s = $1 AND %s = $2 LIMIT %d",
			fmt.Sprintf("%s, %s", _TEST_DB_FIELDNAME_FIELDONE, _TEST_DB_FIELDNAME_FIELDTWO),
			_TEST_DB_TABLENAME,
			"%s",
			"%s",
			limit,
		),
		[]string{_TEST_DB_FIELDNAME_FIELDONE, _TEST_DB_FIELDNAME_FIELDTWO},
		[]interface{}{testValueWhereFieldOne, testValueWhereFieldTwo},
	)
}
