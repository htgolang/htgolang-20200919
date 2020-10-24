package main

import "fmt"

// 见类型知含义 & 可以添加方法
// 现有数据类型
type Counter int

type User map[string]string

type Callback func() error

type Counter2 Counter

func main() {
	var counter Counter //var varname Type
	fmt.Printf("%T, %v\n", counter, counter)
	counter += 10
	fmt.Printf("%T, %v\n", counter, counter)
	var num int = 10
	c2 := Counter(num) + counter //类型转换

	fmt.Printf("%T, %v\n", c2, c2)

	var user User = make(User)

	fmt.Printf("%T, %v\n", user, user)
	user["id"] = "1"
	fmt.Println(user)

	callbacks := map[string]Callback{}

	callbacks["add"] = func() error {
		fmt.Println("add")
		return nil
	}

	callbacks["add"]()

	var c22 Counter2
	fmt.Printf("%T, %v\n", c22, c22)

}
