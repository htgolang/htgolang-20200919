package main

import "fmt"

type User struct {
	id   string
	name string
}

func (user *User) Pfunc() {
	fmt.Println(user.id)
}

func (user User) Vfunc() {
	fmt.Println(user.name)
}

func main() {
	// 对象调用 => 方法值
	var user User
	methodValue01 := user.Pfunc
	methodValue02 := user.Vfunc // 值接收者 =》 user被赋值
	user.id = "xxxx"
	user.name = "yyyy"
	methodValue01() // id?
	methodValue02() // name?
}
