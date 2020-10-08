package main

import (
	"fmt"
)

// 定义Hello
// 无参，无返回值
func sayHello() {
	fmt.Println("hello")
}

// 有参数, 无返回值
func sayHi(name string) {
	fmt.Println("hi, ", name)
}

// 有返回值 add
func add(a int, b int) int {
	fmt.Println(a, b)
	return a + b // return 关键字用来向函数调用者返回结果
}

func main() {
	// 调用函数 函数名称(参数[实参])
	sayHello() // 注意小括号
	sayHello()

	sayHi("kk")
	name := "卫智鹏"
	sayHi(name)

	c := add(2, 1) // 实参 => 形参
	fmt.Println(c)
}
