package main

import "fmt"

func main() {
	height := 1.69
	fmt.Printf("%T, %f, %g, %e\n", height, height, height, height)
	// 运算
	// 算数运算
	// + - * / ++ --
	fmt.Println(1.2 + 1.1)
	fmt.Println(1.2 - 1.1)
	fmt.Println(1.2 * 1.1)
	fmt.Println(1.2 / 1.1)
	height++

	fmt.Println(height)
	height--
	fmt.Println(height)

	// 关系运算
	// > >= <= < [差在某个区间内]
	fmt.Println(1.1 > 1.2)
	fmt.Println(1.1 >= 1.2)
	fmt.Println(1.1 < 1.2)
	fmt.Println(1.1 <= 1.2)
	// 赋值预算
	// =, += , -=, /=, *=
	// a += b a = a + b
	// 类型转换 float64(), float32()

	fmt.Printf("%f %10.3f\n", height, height)

}
