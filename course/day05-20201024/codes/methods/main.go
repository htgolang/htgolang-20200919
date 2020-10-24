package main

import (
	"fmt"
	"methods/models"
)

func main() {
	user := models.NewUser(1, "kk", 32)
	fmt.Println(user)
	// models.AddAge(user) // 修改属性
	user.AddAge()
	fmt.Println(user)
	// fmt.Println(models.GetName(user))
	fmt.Println(user.GetName())

	puser := &user // 取引用 取地址
	// 调用函数如何调用
	(*puser).AddAge() // 解引用 取值
	fmt.Println(*puser)
	fmt.Println((*puser).GetName())

	// GO的语法糖
	puser.AddAge() // 自动解引用 (*puser).AddAge()
	fmt.Println(*puser)
	fmt.Println(puser.GetName()) //自动解引用 //(*puser).GetName()
}
