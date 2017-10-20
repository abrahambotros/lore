package lore

import "testing"

func TestGetPointerSliceLength(t *testing.T) {
	// Test invalid resilience.
	n := 1
	invalids := []interface{}{
		1,
		&n,
		"invalid",
	}
	for _, invalid := range invalids {
		_, err := getPointerSliceLength(invalid)
		if err == nil {
			t.Error("Expect non-nil err since not valid pointer to slice")
			return
		}
	}

	// Test valid lengths.
	var tests = []struct {
		slice  []interface{}
		length uint64
	}{
		{
			slice:  []interface{}{},
			length: 0,
		}, {
			slice:  []interface{}{1},
			length: 1,
		}, {
			slice:  []interface{}{1, "test", struct{}{}},
			length: 3,
		},
	}
	for _, test := range tests {
		l, err := getPointerSliceLength(&test.slice)
		if err != nil {
			t.Error(err)
			return
		}
		if l != test.length {
			t.Errorf("Expect length %d, but got %d", test.length, l)
			return
		}
	}
}
