package main

import (
	"fmt"
	"time"
)

type Addr struct {
	province string
	street   string
	no       string
}

type Tel struct {
	prefix string
	number string
}

// 命名组合(嵌入)
type User struct {
	id       int
	name     string
	addr     Addr
	tel      Tel
	birthday time.Time
}

func main() {
	var user User = User{addr: Addr{province: "陕西省"}}
	fmt.Printf("%T, %#v\n", user, user)

	fmt.Println(user.addr.province)
	user.addr.province = "北京市"
	fmt.Printf("%T, %#v\n", user, user)

	u2 := user
	u2.addr.province = "xxxx"
	fmt.Println(user.addr.province)
}
