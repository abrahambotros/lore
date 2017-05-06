package lore

import "fmt"

/*
testBuildSqlHelper builds the SQL from the Query and compares it against the expected SQL and args.
Returns an error on any errors or deviations from what was expected.

Note that this only can handle a minimal number of args, since Squirrel does not always order args
the same way - thus, trying to test for many args would raise too many permutations of arguments to
easily and quickly test without shifting too focus to accounting for these permutations.
*/
func testBuildSqlHelper(q *Query, expectedSql string, expectedArgs []interface{}) error {
	// Test ToSql.
	qSql, qSqlArgs, err := q.ToSql()
	if err != nil {
		return fmt.Errorf("Error building query to SQL: %s", err.Error())
	}

	// Test if generated SQL is as expected.
	if qSql != expectedSql {
		return fmt.Errorf("Incorrect SQL output: %s. Expected: %s", qSql, expectedSql)
	}

	// Test if args are exactly as expected.
	numQSqlArgs := len(qSqlArgs)
	numExpectedArgs := len(expectedArgs)
	if numQSqlArgs != numExpectedArgs {
		return fmt.Errorf("Incorrect number of SQL args: %d. Expected: %d", numQSqlArgs, numExpectedArgs)
	}
	for i := 0; i < numQSqlArgs; i++ {
		if qSqlArgs[i] != expectedArgs[i] {
			return fmt.Errorf("Incorrect SQL arg at index %d: %+v, expected: %+v", i, qSqlArgs[i], expectedArgs[i])
		}
	}

	// Return with no error.
	return nil
}
