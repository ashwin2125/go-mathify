package arithmetic

// Subtract subtracts two generic numbers together.
func Subtract[T Number](a, b T) T {
	return a - b
}
