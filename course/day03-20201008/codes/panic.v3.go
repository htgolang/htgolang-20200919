package main

import "fmt"

func callback(p bool) {
	if p {
		panic("callback panic")
	}
	fmt.Println("callback running...")
}

func test(p bool) (err error) {
	defer func() {
		if msg := recover(); msg != nil {
			err = fmt.Errorf("%s", msg)
		}
	}()
	callback(p)
	return
}

func main() {
	fmt.Println(test(true))
	fmt.Println(test(false))
	fmt.Println(test(false))
}
