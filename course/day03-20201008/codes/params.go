package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func sayHi(name string) {
	fmt.Println("hi,", name)
}

func add(a int, b int) int {
	return a + b
}

// 参数类型合并
// 参数中多个连续的参数类型相同
// 只保留最后一个类型, 前面连续相同的数据类型可省略
// var a int, b int => var a, b int
func addV2(a, b int) int {
	return a + b
}

func test(a int, b string, c int) {

}

// testV2(a string, b string, c int)
func testV2(a, b string, c int) {

}

// fmt.Println
// append
// 1, 2, 3, 4, 5, ....
// 可变参数
// 定义多个可变参数 不能(只能有一个, 必须定义在形参最后)
func addAll(a, b int, args ...int) int {
	fmt.Println(a, b, args)
	fmt.Printf("%T\n", args)
	// args args切片
	print(args...)
	sum := a + b
	for _, v := range args {
		sum += v
	}
	return sum
}
func print(args ...int) {
	for i, v := range args {
		fmt.Println(i, v)
	}
}

func main() {
	fmt.Println(addV2(3, 4))
	fmt.Println(addAll(1, 2))                      // args ?
	fmt.Println(addAll(1, 2, 3))                   // args ?
	fmt.Println(addAll(1, 2, 3, 4))                // args ?
	fmt.Println(addAll(1, 2, 3, 4, 5, 6, 7, 8, 9)) // args ?

	nums := []int{2, 3, 4, 5}
	fmt.Println(addAll(1, 2, nums...)) // addAll(2, 3, 4, 5)

	print(nums...)
}
