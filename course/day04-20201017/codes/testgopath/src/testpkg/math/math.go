package math

import "mymath"

func Add(n1, n2 int) int {
	mymath.Test()
	return n1 + n2
}

func Sub(n1, n2 int) int {
	return n1 - n2
}

func Mul(n1, n2 int) int {
	return n1 * n2
}

func Div(n1, n2 int) int {
	return n1 / n2
}
