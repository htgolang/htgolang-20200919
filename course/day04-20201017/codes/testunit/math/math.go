package math

func Add(n1, n2 int) int {
	return n1 + n2
}

func Mul(n1, n2 int) int {
	return n1 * n2
}

func Sub(n1, n2 int) int {
	return n1 - n2
}

func Div(n1, n2 int) int {
	if n2 == 0 {
		return 0
	}
	return n1 / n2
}

func Fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * Fact(n-1)
}

func FactV1(n int) int {
	if n == 0 {
		return 1
	}
	rt := 1
	for i := 1; i <= n; i++ {
		rt *= i
	}
	return rt
}
