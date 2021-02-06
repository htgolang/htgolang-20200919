package controllers

import (
	"cmdb/models"
	"cmdb/services"

	"github.com/astaxie/beego"
)

type AuthorizationController struct {
	beego.Controller

	currentUser *models.User
}

func (c *AuthorizationController) Prepare() {
	user := c.GetSession("user")
	if user == nil {
		// 未登录
		c.Redirect("/auth/login", 302)
		return
	}
	if pk, ok := user.(int64); ok {
		if user := services.GetUserById(pk); user != nil {
			c.currentUser = user
		}
	}

	if c.currentUser == nil {
		c.DestroySession()
		c.Redirect("/auth/login", 302)
	} else {
		c.Data["currentUser"] = c.currentUser
	}
}
