package controllers

import (
	"fmt"
	"userMgr/forms"
	"userMgr/services"

	beego "github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

// Login powres user log in
func (c *AuthController) Login() {
	var loginForm = &forms.AuthForm{}
	if user := c.GetSession("user"); user != nil {
		c.Data["user"] = user
		c.Redirect("/user/home/", 302)
		return
	}
	if c.Ctx.Input.IsGet() {
		c.TplName = "user/login.html"
	} else {
		if err := c.ParseForm(loginForm); err != nil {
			panic(err)
		}
		user, err := services.LoginAuth(loginForm)
		if err != nil {
			HandleAuthError(c, err)
		} else {
			c.SetSession("user", user.ID)
			c.Data["user"] = user
			c.Data["form"] = loginForm
			c.Redirect("/user/home", 302)
			fmt.Printf("user %#v logged in.\n", user.Name)
		}
	}
}

func (c *AuthController) Logout() {
	fmt.Println("logout")
	c.DestroySession()
	c.Redirect("/auth/login/", 302)
}
