package arithmetic

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func Add[T Number](a, b T) T {
	return a + b
}
