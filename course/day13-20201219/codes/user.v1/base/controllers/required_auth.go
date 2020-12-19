package controllers

import (
	"github.com/astaxie/beego"
)

type RequiredAuthController struct {
	beego.Controller
}

func (c *RequiredAuthController) Prepare() {
	user := c.GetSession("user")
	if user == nil {
		// 未登录
		c.Redirect("/auth/login", 302)
	}
}
