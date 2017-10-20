package lore

import (
	"testing"

	"github.com/Masterminds/squirrel"
)

/*
TestWhere tests the Where function.

TODO: Add tests for args params.
*/
func TestWhere(t *testing.T) {
	where := Where(map[string]interface{}{
		_TEST_DB_FIELDNAME_FIELD: _TEST_MODEL_FIELD,
	})
	pred, ok := where.pred.(map[string]interface{})
	if !ok {
		t.Error("Unexpected invalid type assertion")
		return
	}
	if pred[_TEST_DB_FIELDNAME_FIELD] != _TEST_MODEL_FIELD {
		t.Errorf(
			"Expected where squirrel.Eq field mapping of %s to %d, but got map: %+v",
			_TEST_DB_FIELDNAME_FIELD,
			_TEST_MODEL_FIELD,
			where.pred,
		)
		return
	}

	// TODO: Add additional tests here for squirrel interfaces.
	where = Where(squirrel.Eq{
		_TEST_DB_FIELDNAME_FIELD: _TEST_MODEL_FIELD,
	})
}
