package main

import (
	"fmt"
)

func main() {
	var (
		年龄     = 32
		weight = 138
	)

	fmt.Println(年龄)

	年龄 = 32 // 赋值，更新变量的值

	年龄, weight = 33, 139

	fmt.Println(年龄, weight)
}
