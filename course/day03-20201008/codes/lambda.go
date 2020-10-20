package main

import (
	"fmt"
	"strings"
)

func print(formatter func(string) string, args ...string) {
	for i, v := range args {
		fmt.Println(i, formatter(v))
	}
}

func add(a, b int) int {
	return a + b
}

func main() {

	// 匿名函数=> 没有名字的函数
	c := func() {
		fmt.Println("我是匿名函数")
	}
	/*
		func c() {
			fmt.Println("我是匿名函数")
		}
	*/

	fmt.Printf("%T\n", c)
	c()
	c()
	c()
	c()

	names := []string{"赵昌建", "kk", "17-赵"}

	star := func(txt string) string {
		return "*" + txt + "*"
	}

	print(star, names...)

	// 1 + 2
	a, b := 1, 2
	fmt.Println(add(a, b))
	fmt.Println(add(1, 2))

	print(func(txt string) string {
		return "|" + txt + "|"
	}, names...)

	fmt.Println(strings.FieldsFunc("AbaabcAbcabc", func(ch rune) bool {
		return ch == 'a'
	}))

	func() {
		fmt.Println("啦啦啦啦啦")
	}()
}
