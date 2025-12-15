package math

func Square(x int) int {
	return x * x
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
