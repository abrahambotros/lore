package lore

import (
	"errors"
	"fmt"
	"reflect"
)

const (
	/*
		STAR provides a hard-coded constant string representing the common star/* character in SQL.
	*/
	STAR string = "*"
	/*
		RETURNING_STAR provides a hard-coded constant string representing the common "RETURNING *"
		clause in SQL.
	*/
	RETURNING_STAR string = "RETURNING *"
	/*
		_ERR_EMPTY_PRIMARY_KEY provides a common error message used by this package to indicate an
		empty primary key is found on a ModelInterface instance when a non-empty one is required.
	*/
	_ERR_EMPTY_PRIMARY_KEY string = "Empty primary key encountered; cannot perform functions requiring primary keys"
)

/*
isPointer returns true if the given interface is a pointer; false otherwise.
*/
func isPointer(i interface{}) bool {
	return reflect.ValueOf(i).Kind() == reflect.Ptr
}

/*
isPointerToSlice returns true if the given interface is a pointer to a slice; false otherwise.
*/
func isPointerToSlice(i interface{}) bool {
	return isPointer(i) && reflect.TypeOf(reflect.ValueOf(i).Elem().Interface()).Kind() == reflect.Slice
}

/*
getPointerSliceLength returns the length of the slice pointed to by the given pointer. If the
underlying object is not a slice, an error is returned.
*/
func getPointerSliceLength(i interface{}) (uint64, error) {
	if isPointerToSlice(i) {
		return uint64(reflect.ValueOf(i).Elem().Len()), nil
	}

	// If reach here, then i was not a pointer to a slice; return error.
	return 0, fmt.Errorf("Input was not a valid pointer to slice: %+v", i)
}

/*
getMapKeysVals takes in a map and returns two correspondingly-ordered slices - one containing the
string keys of the map, and the other containing the corresponding values.
*/
func getMapKeysVals(m map[string]interface{}) (keys []string, vals []interface{}, err error) {
	// If invalid map, return error now.
	if m == nil {
		return nil, nil, errors.New("Input map cannot be nil")
	}

	// Create keys and vals slices with appropriate length/capacity.
	n := len(m)
	keys = make([]string, n)
	vals = make([]interface{}, n)

	// Iterate through map keys/vals and place into slices.
	i := 0
	for key, val := range m {
		keys[i] = key
		vals[i] = val
		i++
	}

	// Return final keys/vals with success.
	return keys, vals, nil
}
