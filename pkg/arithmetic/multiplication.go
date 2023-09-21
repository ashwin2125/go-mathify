package arithmetic

// Multiply multiplies two generic numbers together.
func Multiply[T Number](a, b T) T {
	return a * b
}