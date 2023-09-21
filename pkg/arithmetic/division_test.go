package arithmetic

import (
	"reflect"
	"testing"
)

func TestDivide(t *testing.T) {
	testCases := []struct {
		a      interface{}
		b      interface{}
		result interface{}
		err    error
	}{
		{10, 2, int(5), nil},
		{int64(10), int64(2), int64(5), nil},
		{10.0, 2.0, 5.0, nil},
		{10, 0, nil, ErrDivisionByZero},
		{0, 0, nil, ErrDivisionByZero},
		{"10", 2, nil, ErrUnsupportedType},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)

		if err != nil && err != tc.err {
			t.Errorf("Divide(%v, %v) returned an unexpected error: got %v; want %v", tc.a, tc.b, err, tc.err)
			continue
		}

		if !reflect.DeepEqual(result, tc.result) {
			t.Errorf("Divide(%v, %v) = %v; want %v", tc.a, tc.b, result, tc.result)
		}
	}
}
