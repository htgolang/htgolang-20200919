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
	prefix string
	number string
	// province string
}

// func (tel *Tel) SetProvince(name string) {
// 	fmt.Println("tel.setprovince")
// }

// 命名组合方式
type User struct {
	id   int
	name string
	Addr
	Tel
	birthday time.Time
	province string
}

// func (user *User) SetProvince(name string) {
// 	fmt.Println("user.setprovince")
// }

func main() {
	var user User
	fmt.Printf("%#v\n", user)
	// user.Addr.SetProvince("陕西省")
	user.SetProvince("陕西省")
	fmt.Printf("%#v\n", user)
}
