package arithmetic

import (
	"errors"
	"math"
	"reflect"
)

// Custom errors
var ErrDivisionByZero = errors.New("division by zero")
var ErrUnsupportedType = errors.New("unsupported type")

// Divide function accepts two interfaces and returns an interface along with an error.
func Divide(a, b interface{}) (interface{}, error) {
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	// Check for type consistency between 'a' and 'b'
	if valA.Kind() != valB.Kind() {
		return nil, ErrUnsupportedType
	}

	switch valA.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ai := valA.Int()
		bi := valB.Int()

		if bi == 0 {
			return nil, ErrDivisionByZero
		}

		switch a.(type) {
		case int:
			return int(ai / bi), nil
		case int64:
			return ai / bi, nil
		}

	case reflect.Float32, reflect.Float64:
		af := valA.Float()
		bf := valB.Float()

		if bf == 0 {
			return math.NaN(), nil
		}

		return af / bf, nil
	default:
		return nil, ErrUnsupportedType
	}

	return nil, errors.New("unreachable code")
}
