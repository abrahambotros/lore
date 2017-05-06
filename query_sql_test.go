package lore

import (
	"fmt"
	"testing"
)

/*
testBuildSqlHelper builds the SQL from the Query and compares it against the expected SQL and args.
Note that we return an error instead of going through the testing module so that callers can try
different permutations for SQL args before actually failing a test.
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

func testBuildSqlHelperWithArgPermutations(t *testing.T, q *Query, expectedSqlBase string, argNames []string, argValues []interface{}) {
	numArgs := len(argNames)
	if numArgs != len(argValues) {
		t.Errorf("argNames and argValues must have same length. argNames:%+v, argValues:%+v", argNames, argValues)
		return
	}

	// If more than 2 args, fatal err that unsupported.
	if numArgs > 2 {
		t.Errorf("Maximum of 2 args currently supported, got %d in %+v", numArgs, argNames)
		return
	}

	var expectedSql string
	var expectedArgs []interface{}
	var err error
	for i := 0; i < numArgs; i++ {
		if i == 0 {
			// Forward order.
			expectedSql = fmt.Sprintf(expectedSqlBase, argNames[0], argNames[1])
			expectedArgs = argValues
		} else {
			// Reverse order.
			expectedSql = fmt.Sprintf(expectedSqlBase, argNames[1], argNames[0])
			expectedArgs = []interface{}{argValues[1], argValues[0]}
		}
		err = testBuildSqlHelper(q, expectedSql, expectedArgs)
		if err == nil {
			// If no error, then break immediately.
			break
		}
	}
	// If still err here, then all permutations failed.
	if err != nil {
		t.Errorf(
			"testBuildSql failed all SQL permutations with expectedSqlBase: %s, argNames: %+v, argValues: %+v. Last err: %s",
			expectedSqlBase, argNames, argValues, err.Error(),
		)
		return
	}
}
