package gosample

import "errors"

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("divide by zero.")
	}

	answer := x / y
	rest := x % y

	return answer, rest, nil
}
