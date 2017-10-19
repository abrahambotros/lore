package lore

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
)

/*
TestSelectModelByPrimaryKey tests the SelectModelByPrimaryKey function.
*/
func TestSelectModelByPrimaryKey(t *testing.T) {
	// Test invalid ModelInterface instance.
	tmi := createTestModelInvalidInstance()

	// Test empty primary key resilience.
	pk := tmi.DbPrimaryFieldKey()
	if pk != "" {
		t.Error("Expected empty primary field key for invalid instance")
		return
	}
	found, err := SelectModelByPrimaryKey(tmi, nil, nil)
	if found != false || err == nil {
		t.Error("Expected not found and non-nil err for invalid instance with empty primary key")
		return
	}

	tm := createTestModelInstance()
	db, dbMock := getTestSqlxDb(t)
	dbMock.ExpectQuery(fmt.Sprintf(
		"^SELECT \\* FROM %s WHERE %s =", _TEST_DB_TABLENAME, _TEST_DB_FIELDNAME_ID,
	)).WithArgs(_TEST_MODEL_ID)
	SelectModelByPrimaryKey(tm, db, newTestModelEmpty())
	err = dbMock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
		return
	}
}

/*
TestSelectModelWhere tests the SelectModelWhere function.
*/
func TestSelectModelWhere(t *testing.T) {
	tm := createTestModelInstance()
	db, dbMock := getTestSqlxDb(t)
	dbMock.ExpectQuery(fmt.Sprintf(
		"^SELECT \\* FROM %s WHERE %s", _TEST_DB_TABLENAME, _TEST_DB_FIELDNAME_FIELD,
	)).WithArgs(_TEST_MODEL_FIELD)
	SelectModelWhere(tm, db, newSqlPart(squirrel.Eq{
		_TEST_DB_FIELDNAME_FIELD: _TEST_MODEL_FIELD,
	}), newTestModelEmpty())
	err := dbMock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
		return
	}
}

/*
TestSelectModelsWhere tests the SelectModelsWhere function.
*/
func TestSelectModelsWhere(t *testing.T) {
	tm := createTestModelInstance()
	db, dbMock := getTestSqlxDb(t)
	dbMock.ExpectQuery(fmt.Sprintf(
		"^SELECT \\* FROM %s WHERE %s.*LIMIT 3", _TEST_DB_TABLENAME, _TEST_DB_FIELDNAME_FIELD,
	)).WithArgs(_TEST_MODEL_FIELD)
	var limit uint64
	limit = 3
	SelectModelsWhere(tm, db, newSqlPart(squirrel.Eq{
		_TEST_DB_FIELDNAME_FIELD: _TEST_MODEL_FIELD,
	}), &limit, newTestModelEmptyList())
	err := dbMock.ExpectationsWereMet()
	if err != nil {
		t.Error(err)
		return
	}
}
