package main

import (
	"fmt"
	"time"
)

type User struct {
	id       int
	name     string
	addr     string
	tel      string
	birthday time.Time
}

type Addr struct {
}

// New函数
func NewUser(id int, name, addr, tel string, birthday time.Time) *User {
	return &User{id, name, addr, tel, birthday}
}

func main() {
	var user = &User{id: 10, name: "kk"}
	fmt.Printf("%T, %#v\n", user, user)

	// new函数初始化
	user = new(User)
	fmt.Printf("%T, %#v\n", user, user)

	u2 := new(User)
	fmt.Printf("%T, %#v\n", u2, u2)

	//属性访问
	fmt.Println(u2.name)
	u2.name = "kk"
	u2.id = 100

	fmt.Printf("%T, %#v\n", u2, u2)

	u3 := NewUser(1, "kk", "", "", time.Now())
	fmt.Printf("%T, %#v\n", u3, u3)
}
