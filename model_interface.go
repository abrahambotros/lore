package lore

// TODO: CLEANUP
// /*
// Model provides a base model definition.
// */
// type Model struct{}

/*
ModelInterface provides a generic interface for enabling external models to interface with internal
machinery here.
*/
type ModelInterface interface {
	/*
		DbTableName provides the name of the corresponding table in the database for this model.
	*/
	DbTableName() string
	/*
		DbFieldMap builds a map for this model instance from field/column name to this instance's
		current value for that field/column. Note that this should NOT include auto-managed db
		fields, such as SERIAL/AUTO-INCREMENTING KEYS, unless you want to try to specifically
		override these each time (typically not the case).

		This is used to provide Insert and Update functionality convenience by providing all of the
		columns and values to write into the table row.
	*/
	DbFieldMap() map[string]interface{}
	/*
		DbPrimaryFieldKey returns a string indicating the column name of the primary field of the
		model. This is used for Update and Delete queries for the WHERE condition.

		If the implementing model does not have such a field, an empty string can be returned, which
		will result in an error being thrown if any methods are called that require a non-empty
		field (such as an UPDATE that should have a non-empty WHERE condition for updating by
		primary key). If multiple columns define the primary field, you will have to implement the
		condition yourself instead.
	*/
	DbPrimaryFieldKey() string
	/*
		DbPrimaryFieldValue returns the current value of the primary field of the model. This is
		used for Update and Delete queries for the WHERE condition in conjunction with
		DbPrimaryFieldKey.

		If the implementing model does not have such a primary field, nil can be returned. Note that
		an error will be thrown if the DbPrimaryFieldKey method returns an empty string (see
		DbPrimaryFieldKey), but an error is not necessarily thrown if DbPrimaryFieldKey is
		non-empty and DbPrimaryFieldValue is nil (i.e., nil is a valid value to set for a primary
		field value if you really want).
	*/
	DbPrimaryFieldValue() interface{}
}

/**
 * TODO: JUMP
 * - In README, instruct to send result as in sqlx...
 */

// TODO: REMOVE
// /*
// dbColumnValueMap parses the Model's fields and tags to build a map from column (name) to value for
// columns that are marked for LORE consumption.
// */
// func (m *Model) dbColumnValueMap() map[string]interface{} {
// 	// Parse model.
// 	mValue := reflect.ValueOf(m)
// 	mType := mValue.Type()
// 	mNumFields := mValue.NumField()

// 	// Init return map.
// 	cvm := map[string]interface{}{}

// 	// Iterate model fields.
// 	for i := 0; i < mNumFields; i++ {
// 		// TODO: Use fatih/structs
// 		mType.Field(i).Name
// 		mValue.Field(i).Interface()
// 	}
// }
