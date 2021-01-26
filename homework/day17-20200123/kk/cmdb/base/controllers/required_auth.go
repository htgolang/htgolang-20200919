package controllers

import (
	"cmdb/services"

	"github.com/astaxie/beego"
)

type RequiredAuthController struct {
	beego.Controller
}

func (c *RequiredAuthController) Prepare() {

	c.Data["currentUser"] = nil
	c.Data["navKey"] = ""

	user := c.GetSession("user")
	if user == nil {
		// 未登录
		c.Redirect("/auth/login", 302)
		return
	}
	if pk, ok := user.(int64); ok {
		if user := services.GetUserById(pk); user != nil {
			c.Data["currentUser"] = user
		}
	}

	if c.Data["currentUser"] == nil {
		c.DestroySession()
		// 未登录
		c.Redirect("/auth/login", 302)
		return
	}
}
