## The Generic Division Problem in Go

Go is a strongly typed language, and it doesn't provide generics (at least until Go 1.18). This makes it challenging to write a single function that can handle division for multiple types.

### Approaches Considered

#### Using Go 1.18 Generics

Go 1.18 introduces generics, but this feature doesn't fully cater to arithmetic operations on generic types. It lacks native operations for generic types, meaning we would still need to cast types, defeating the purpose of using generics.

#### Using Empty Interface and Type Assertion

Another approach is to use empty interfaces (`interface{}`) for accepting all types and then use type assertions to determine the type at runtime. The problem with this approach is it requires a large switch-case construct to handle each type, making the code cumbersome and hard to maintain.

#### Custom Structs for Numbers

We considered defining a custom struct that can hold any numeric type and implements methods for arithmetic operations. The issue with this is the overhead of managing custom structs and losing native type benefits, such as immediate compatibility with other libraries.

### Final Approach: Using Reflection

We have chosen to use Go's `reflect` package to implement a generic division function, which examines the types at runtime. This allows us to perform division while handling type-related corner cases.

```go
// Custom errors
var ErrDivisionByZero = errors.New("division by zero")
var ErrUnsupportedType = errors.New("unsupported type")

// Divide function accepts two interfaces and returns an interface along with an error.
func Divide(a, b interface{}) (interface{}, error) {
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

```

## Shortcomings and Workarounds

- **Type Mismatch**: Our `Divide` function checks for type consistency between 'a' and 'b' and returns an error if there's a mismatch. This is a limitation but also a feature to prevent accidental operations between incompatible types.
- **Zero Division**: The function returns a custom error for zero divisions for integer types and `NaN` for floating-point types.

## Running Tests

To run tests for the `arithmetic` package, navigate to the package directory and execute:

```bash
go test -v
```

## Contributions

Feel free to submit PRs or to open issues.
```
