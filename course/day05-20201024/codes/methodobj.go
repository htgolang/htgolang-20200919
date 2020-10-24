package main

import "fmt"

type User struct {
}

func (user *User) Pfunc() {
	fmt.Println("pfunc")
}

func (user User) Vfunc() {
	fmt.Println("vfunc")
}

func main() {
	// 对象调用 => 方法值
	var user User
	methodValue01 := user.Pfunc
	fmt.Printf("%T\n", methodValue01)
	methodValue01()

	methodValue02 := user.Vfunc
	fmt.Printf("%T\n", methodValue02)
	methodValue02()

	u2 := &user
	methodValue11 := u2.Pfunc

	fmt.Printf("%T\n", methodValue11)
	methodValue11()

	methodValue12 := u2.Vfunc
	fmt.Printf("%T\n", methodValue12)
	methodValue12()

	// 结构体名称 => 方法表达式

	methodValue21 := User.Vfunc
	// methodValue22 := User.Pfunc // 错误
	fmt.Printf("%T\n", methodValue21)
	// fmt.Printf("%T\n", methodValue22)
	methodValue21(user)
	// methodValue21(u2)

	methodValue31 := (*User).Pfunc
	methodValue32 := (*User).Vfunc
	fmt.Printf("%T\n", methodValue31)
	fmt.Printf("%T\n", methodValue32)
	methodValue31(u2)
	methodValue32(u2)

	// 定义接收者方法
	// 值接收者 => 自动生成指针接收者 => 值、指针的
	// 指针接收者 => 只有指针接收者方法

}
