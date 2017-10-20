package lore

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
)

func TestNewSqlPart(t *testing.T) {
	newSqlPart(squirrel.Eq{
		_TEST_DB_FIELDNAME_FIELD: _TEST_MODEL_FIELD,
	})
	newSqlPart(squirrel.Lt{
		_TEST_DB_FIELDNAME_FIELD: _TEST_MODEL_FIELD,
	})
	newSqlPart(map[string]interface{}{
		_TEST_DB_FIELDNAME_FIELD: _TEST_MODEL_FIELD,
	})
}

func TestSqlPartToSql(t *testing.T) {
	// Test invalid input resilience.
	sp := newSqlPart(10)
	sql, args, err := sp.ToSql()
	if err == nil {
		t.Error("Expect non-nil err since invalid pred")
		return
	}

	// Test squirrel interface.
	sp = newSqlPart(squirrel.Eq{
		_TEST_DB_FIELDNAME_FIELD: _TEST_MODEL_FIELD,
	})
	sql, args, err = sp.ToSql()
	if sql != fmt.Sprintf("%s = ?", _TEST_DB_FIELDNAME_FIELD) {
		t.Errorf("Invalid sql generated: %s", sql)
		return
	}
	arg0, ok := args[0].(int64)
	if !ok {
		t.Error("Unexpected error casting to int64")
		return
	}
	if arg0 != _TEST_MODEL_FIELD {
		t.Errorf("Expected %d, got %d", _TEST_MODEL_FIELD, arg0)
		return
	}

	// Test string.
	sp = newSqlPart(
		fmt.Sprintf("%s = ?", _TEST_DB_FIELDNAME_FIELD),
		_TEST_MODEL_FIELD,
	)
	sql, args, err = sp.ToSql()
	if sql != fmt.Sprintf("%s = ?", _TEST_DB_FIELDNAME_FIELD) {
		t.Errorf("Invalid sql generated: %s", sql)
	}
	arg0, ok = args[0].(int64)
	if !ok {
		t.Error("Unexpected error casting to int64")
		return
	}
	if arg0 != _TEST_MODEL_FIELD {
		t.Errorf("Expected %d, got %d", _TEST_MODEL_FIELD, arg0)
		return
	}
}
