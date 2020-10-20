package main

import "fmt"

func callback(p bool) {
	if p {
		panic("callback panic")
	}
	fmt.Println("callback running...")
}

func test(p bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出错了, 进行恢复", err)
		}
	}()
	callback(p)
}

func main() {
	test(true)
	test(false)
	test(false)
}
