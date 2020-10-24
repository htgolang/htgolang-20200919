package main

import (
	"fmt"
	"methods/models"
)

func main() {
	user := models.NewUser(1, "kk", 32)
	fmt.Println(user)
	// models.AddAge(user) // 修改属性
	user.AddAge() // 结构的对象.方法名称()
	fmt.Println(user)
	// fmt.Println(models.GetName(user))
	fmt.Println(user.GetName())

	u2 := *user
	fmt.Printf("%T, %#v\n", u2, u2)

	(&u2).AddAge()
	fmt.Println(u2)
	fmt.Println((&u2).GetName())
	// Go语法糖
	u2.AddAge() // (&u2).AddAge()
	fmt.Println(u2)
	fmt.Println(u2.GetName()) // (&u2).GetName()
}
