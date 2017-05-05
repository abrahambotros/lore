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
ExecuteThenParseSingle ...

TODO: Note that handle sql.ErrNoRows and return nil instead.
*/
func (q *Query) ExecuteThenParseSingle(db *sqlx.DB, resultSingle interface{}) (numRowsAffected int, err error) {
	return q.execute(db, resultSingle, _EXECUTE_MODE_PARSE_SINGLE)
}

/*
ExecuteThenParseList ...

TODO: Note resultList should be POINTER to list.

TODO: Note warning of Select vs Queryx - should be bounded.
*/
func (q *Query) ExecuteThenParseList(db *sqlx.DB, resultList interface{}) (numRowsAffected int, err error) {
	return q.execute(db, resultList, _EXECUTE_MODE_PARSE_LIST)
}

/*
execute ...

TODO: Sets result to nil if
*/
func (q *Query) execute(db *sqlx.DB, result interface{}, mode executeMode) (numRowsAffected int, err error) {
	// Handle invalid input.
	if db == nil {
		return 0, errors.New("db cannot be nil")
	}
	if mode != _EXECUTE_MODE_NO_PARSE && mode != _EXECUTE_MODE_PARSE_SINGLE && mode != _EXECUTE_MODE_PARSE_LIST {
		return 0, fmt.Errorf("Invalid execute mode: %d", mode)
	}

	// Get SQL.
	qSql, qSqlArgs, err := q.ToSql()
	if err != nil {
		return 0, fmt.Errorf("Error building Query SQL: %s", err.Error())
	}

	// Execute SQL with args on db.
	if mode == _EXECUTE_MODE_NO_PARSE {
		_, err = db.Exec(qSql, qSqlArgs...)
	} else if mode == _EXECUTE_MODE_PARSE_SINGLE {
		err = db.Get(result, qSql, qSqlArgs...)
	} else if mode == _EXECUTE_MODE_PARSE_LIST {
		err = db.Select(result, qSql, qSqlArgs...)
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
	if mode == _EXECUTE_MODE_PARSE_SINGLE {
		// If single mode, then assuming sql.ErrNoRows was caught above, we must have retrieved a single
		// result.
		numRowsAffected = 1
	} else if mode == _EXECUTE_MODE_PARSE_LIST {
		// Otherwise, if in multi mode, then explicitly count number of items returned via
		// reflection, assuming the input was indeed a pointer to a slice.
		numRowsAffected, err = getPointerSliceLength(result)
		if err != nil {
			return 0, err
		}
	}

	// Return with success.
	return numRowsAffected, nil
}
