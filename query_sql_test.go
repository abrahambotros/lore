package lore

import "testing"

func testBuildSqlHelper(t *testing.T, q *Query, expectedSql string, expectedArgs []interface{}) {
	// Test ToSql.
	qSql, qSqlArgs, err := q.ToSql()
	if err != nil {
		t.Fatalf("Error building query to SQL: %s", err.Error())
	}

	// Test if generated SQL is as expected.
	if qSql != expectedSql {
		t.Fatalf("Incorrect SQL output: %s. Expected: %s", qSql, expectedSql)
	}

	// Test if args are exactly as expected.
	numQSqlArgs := len(qSqlArgs)
	numExpectedArgs := len(expectedArgs)
	if numQSqlArgs != numExpectedArgs {
		t.Fatalf("Incorrect number of SQL args: %d. Expected: %d", numQSqlArgs, numExpectedArgs)
		return
	}
	for i := 0; i < numQSqlArgs; i++ {
		if qSqlArgs[i] != expectedArgs[i] {
			t.Fatalf("Incorrect SQL arg at index %d: %+v, expected: %+v", i, qSqlArgs[i], expectedArgs[i])
		}
	}
}
