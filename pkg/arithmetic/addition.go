package arithmetic

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two generic numbers together.
func Add[T Number](a, b T) T {
	return a + b
}
