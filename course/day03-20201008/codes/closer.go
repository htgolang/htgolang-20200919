package main

import "fmt"

// 生成一个函数
// add(int) int
// left + base
// 变量的生命周期 => 内存中存在的时间
func addBase(base int) func(int) int {
	return func(num int) int {
		return base + num
	}
}

func main() {
	add2 := addBase(2)
	fmt.Printf("%T\n", add2)
	fmt.Println(add2(10))
}
