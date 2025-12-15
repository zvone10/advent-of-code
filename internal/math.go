package internal

func Square(x int) int {
	return x * x
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}

	return -x
}
