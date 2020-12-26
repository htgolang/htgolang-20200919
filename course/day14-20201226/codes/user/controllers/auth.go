package controllers

import (
	"fmt"
	"user/errors"
	"user/forms"
	"user/services"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	if user := c.GetSession("user"); user != nil {
		c.Redirect("/user/query", 302)
		return
	}
	// 打开
	// 点击登陆
	form := &forms.LoginForm{}
	// 错误信息 attr = []error
	errors := errors.NewErrors()

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			if user := services.Auth(form); user != nil {
				fmt.Println("登陆成功")
				c.SetSession("user", user.ID)
				c.Redirect("/user/query", 302)
				return
			} else {
				errors.AddError("default", "用户名或密码错误")
				errors.AddError("username", "xxxx")
			}
		}
	}
	c.Data["form"] = form
	c.Data["errors"] = errors
	c.TplName = "auth/login.html"
}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Redirect("/auth/login/", 302)
}
