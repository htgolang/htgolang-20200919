package main

import "fmt"

func factV1(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 1
	} else {
		rt := 1
		for i := 1; i <= n; i++ {
			rt *= i
		}
		return rt
	}
}

func fact(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	/*
		阶乘
		n! = 1 * 2 * 3 * ... * n
		=> (1 * ... * n-1) *n
		=> (n-1)! * n

		0! = 1
		fact(n) = n * fact(n-1)
		递归=> 结束条件
		fact(3) = 3 * fact(2)
		fact(2) = 2 * fact(1)
		fact(1) = 1 * fact(0)
		fact(0) = 1
	*/
	fmt.Println(fact(5))
}
