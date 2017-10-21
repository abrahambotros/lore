package lore

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type executeMode int8

const (
	_EXECUTE_MODE_NO_PARSE     executeMode = 0
	_EXECUTE_MODE_PARSE_SINGLE executeMode = 1
	_EXECUTE_MODE_PARSE_LIST   executeMode = 2
)

/*
Execute wraps sqlx.DB.Exec, and does not parse the result into any struct. However, it still returns
the basic info of numRowsAffected IN SOME CASES (see below).

NOTE: As mentioned at http://jmoiron.github.io/sqlx/#exec access to numRowsAffected (via
sqlx.Result.RowsAffected) is db-driver-dependent.
*/
func (q *Query) Execute(db *sqlx.DB) (numRowsAffected uint64, err error) {
	return q.execute(db, nil, _EXECUTE_MODE_NO_PARSE)
}

/*
ExecuteThenParseSingle wraps sqlx.DB.Get to execute the query and then parse the result into the
given single struct in resultSinglePtr. Scanning the db results into the result interface is done
according to typical sqlx scanning behavior for sqlx.DB.Get.

When sql.ErrNoRows is encountered by the underlying sqlx.DB.Get query (any time no matching row is
found for this query), this function returns with found equal to false (naturally), but DOES NOT
RETURN ANY ERROR (even though the sql package did).

Note that resultSinglePtr should be a POINTER TO A SINGLE STRUCT that you want to scan the results
into. An error will be returned if resultSinglePtr is not detected as a pointer to a single struct.
*/
func (q *Query) ExecuteThenParseSingle(db *sqlx.DB, resultSinglePtr interface{}) (found bool, err error) {
	numRowsAffected, err := q.execute(db, resultSinglePtr, _EXECUTE_MODE_PARSE_SINGLE)
	return (numRowsAffected > 0), err
}

/*
ExecuteThenParseList wraps sqlx.DB.Select to execute the query and then parse the result into the
given list of structs in resultSinglePtr. Scanning the db results into the result interface is done
according to typical sqlx scanning behavior for sqlx.DB.Select.

Note that resultListPtr should be a POINTER TO A LIST OF STRUCTS that you want to scan the results
into. An error will be returned if resultListPtr is not detected as a pointer to a list.

This function returns numRowsAffected equal to the length derived from the result slice after the
query has been run.

NOTE: Any callers of this function should ensure the underlying query is appropriately bounded, as
sqlx.DB.Select will load the entire result set into memory at once; see
http://jmoiron.github.io/sqlx/#getAndSelect for more information.
*/
func (q *Query) ExecuteThenParseList(db *sqlx.DB, resultListPtr interface{}) (numRowsAffected uint64, err error) {
	return q.execute(db, resultListPtr, _EXECUTE_MODE_PARSE_LIST)
}

/*
execute provides the underlying execute functionality for executing SQL on a sqlx.DB. When mode is
NO_PARSE, just calls db.Exec; when PARSE_SINGLE, calls db.Get; when PARSE_LIST, calls db.Select.

Note that in the PARSE_SINGLE case, when sql.ErrNoRows is encountered by the underlying sqlx.DB.Get
query (any time no matching row is found for this query), this function returns 0 rows affected
(naturally), but DOES NOT RETURN ANY ERROR.
*/
func (q *Query) execute(db *sqlx.DB, resultPtr interface{}, mode executeMode) (numRowsAffected uint64, err error) {
	// Handle invalid input.
	if db == nil {
		return 0, errors.New("db cannot be nil")
	}
	switch mode {
	case _EXECUTE_MODE_NO_PARSE:
		break
	case _EXECUTE_MODE_PARSE_SINGLE:
		if !isPointer(resultPtr) || isPointerToSlice(resultPtr) {
			return 0, fmt.Errorf("Result pointer cannot point to slice when in ParseSingle mode. ResultPtr: %+v", resultPtr)
		}
	case _EXECUTE_MODE_PARSE_LIST:
		if !isPointer(resultPtr) || !isPointerToSlice(resultPtr) {
			return 0, fmt.Errorf("Result pointer must point to slice when in ParseList mode. ResultPtr: %+v", resultPtr)
		}
	default:
		return 0, fmt.Errorf("Invalid execute mode: %d", mode)
	}

	// Get SQL.
	qSql, qSqlArgs, err := q.ToSql()
	if err != nil {
		return 0, fmt.Errorf("Error building Query SQL: %s", err.Error())
	}

	// Execute SQL with args on db.
	var rawExecResult sql.Result
	switch mode {
	case _EXECUTE_MODE_NO_PARSE:
		rawExecResult, err = db.Exec(qSql, qSqlArgs...)
	case _EXECUTE_MODE_PARSE_SINGLE:
		err = db.Get(resultPtr, qSql, qSqlArgs...)
	case _EXECUTE_MODE_PARSE_LIST:
		err = db.Select(resultPtr, qSql, qSqlArgs...)
	}

	// Handle any specific soft errors.
	if err == sql.ErrNoRows {
		// If specifically encountered this error, then set the number of rows affected to 0, and
		// return with no error immediately.
		return 0, nil
	}

	// Handle hard errors.
	if err != nil {
		return 0, fmt.Errorf("Error running Query: %s. SQL:%s, SQL args:%+v", err.Error(), qSql, qSqlArgs)
	}

	// Do any additional per-mode handling of result/return values.
	switch mode {
	case _EXECUTE_MODE_NO_PARSE:
		// If no-parse mode, then attempt using the raw sql.Result to determine number of rows
		// affected
		numRowsAffectedInt, err := rawExecResult.RowsAffected()
		if err != nil {
			return 0, err
		}
		numRowsAffected = uint64(numRowsAffectedInt)
	case _EXECUTE_MODE_PARSE_SINGLE:
		// If single mode, then assuming sql.ErrNoRows was caught above, we must have retrieved a
		// single result.
		numRowsAffected = 1
	case _EXECUTE_MODE_PARSE_LIST:
		// Otherwise, if in multi mode, then explicitly count number of items returned via
		// reflection, assuming the input was indeed a pointer to a slice.
		numRowsAffected, err = getPointerSliceLength(resultPtr)
		if err != nil {
			return 0, err
		}
	}

	// Return with success.
	return numRowsAffected, nil
}
