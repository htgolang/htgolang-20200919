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

func (addr *Addr) SetProvince(province string) {
	addr.province = province
}

type Tel struct {
	prefix   string
	number   string
	province string
}

// 命名组合方式
type User struct {
	id       int
	name     string
	addr     Addr
	tel      Tel
	birthday time.Time
	province string
}

func main() {
	var user User
	fmt.Printf("%#v\n", user)
	user.addr.SetProvince("陕西省")
	fmt.Printf("%#v\n", user)
}
