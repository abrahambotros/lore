package lore

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
)

func TestBuildWhereColumnEqualsSubquery(t *testing.T) {
	fieldTestValueStart := "testStart"
	fieldTestValueEnd := "testEnd"

	// Build subquery.
	qSub := NewQuery(newTestModelEmpty())
	qSub.SetSqlBuilder(
		BuildWhereColumnEqualsSubquerySelect(
			_TEST_DB_FIELDNAME_ID,
			qSub.BuildSqlSelect(_TEST_DB_FIELDNAME_ID).
				Where(squirrel.Eq{
					_TEST_DB_FIELDNAME_FIELD: fieldTestValueStart,
				}),
		),
	)

	// // Build subquery where wrapper.
	// whereQSubResult, err := BuildWhereColumnEqualsSubquery(
	// 	_TEST_DB_FIELDNAME_ID,
	// 	qSub,
	// )
	// if err != nil {
	// 	t.Errorf("Unexpected error in BuildWhereColumnEqualsSubquery: %s", err.Error())
	// 	return
	// }

	// Build full outside query.
	q := NewQuery(newTestModelEmpty())
	q.SetSqlBuilder(
		q.BuildSqlUpdate().
			Set(_TEST_DB_FIELDNAME_FIELD, fieldTestValueEnd).
			Where(qSub),
	)

	// Test SQL and args.
	qSql, qArgs, err := q.ToSql()
	if err != nil {
		t.Errorf("Error building query to SQL and args: %s", err.Error())
		return
	}

	expectedSql := fmt.Sprintf(
		"UPDATE %s SET %s = $1 WHERE %s = ( SELECT %s FROM %s WHERE %s = $2 );",
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELD,
		_TEST_DB_FIELDNAME_ID,
		_TEST_DB_FIELDNAME_ID,
		_TEST_DB_TABLENAME,
		_TEST_DB_FIELDNAME_FIELD,
	)
	if qSql != expectedSql {
		t.Errorf("Expected SQL: %s\nGot: %s", expectedSql, qSql)
		return
	}

	expectedArgs := []interface{}{fieldTestValueEnd, fieldTestValueStart}
	if len(qArgs) != len(expectedArgs) {
		t.Errorf("Invalid number of args. Expected args: %+v\nGot: %+v", expectedArgs, qArgs)
		return
	}
	for i, qArg := range qArgs {
		if qArg != expectedArgs[i] {
			t.Errorf("Expected args: %+v\nGot: %+v\nInequality found at index %d", expectedArgs, qArgs, i)
			return
		}
	}
}
