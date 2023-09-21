package arithmetic

import "errors"

func Divide[T Number](a, b T) (T, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
