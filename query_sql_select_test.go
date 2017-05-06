package lore

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
)

func TestBuildSqlSelect(t *testing.T) {
	SetConfigDefault()
	tm := createTestModelInstance()
	testValueWhereField := 2
	limit := uint64(10)

	// Build query and SQL.
	q := NewQuery(tm)
	q.SetSqlBuilder(
		q.BuildSqlSelect(_TEST_DB_FIELDNAME_ID).
			Where(squirrel.Eq{
				_TEST_DB_FIELDNAME_FIELD: testValueWhereField,
			}).
			Limit(limit),
	)

	// Delegate to SQL test helper.
	err := testBuildSqlHelper(
		q,
		fmt.Sprintf(
			"SELECT %s FROM %s WHERE %s = $1 LIMIT %d",
			_TEST_DB_FIELDNAME_ID,
			_TEST_DB_TABLENAME,
			_TEST_DB_FIELDNAME_FIELD,
			limit,
		),
		[]interface{}{testValueWhereField},
	)
	if err != nil {
		t.Error(err)
		return
	}
}
