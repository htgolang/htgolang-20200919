/*
	用户未认证钩子函数
*/

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RequireAuth struct {
	beego.Controller
}

func (c *RequireAuth) Prepare() {
	// 用户未登录时重定向至登录页面
	user := c.GetSession("user")
	fmt.Println("获得用户session：", user)
	if user == nil {
		c.Redirect("/auth/login", 302)
	}
}
