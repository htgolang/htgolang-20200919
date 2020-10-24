package main

import (
	"fmt"
	"time"
)

// 指针，地址，切片，映射
type User struct {
	id       int
	name     string
	addr     string
	tel      string
	birthday time.Time
}

func main() {
	user := User{}
	u2 := user
	u2.name = "kk"

	fmt.Printf("%#v\n", u2)
	fmt.Printf("%#v\n", user)
}
