package controller

import (
	"fmt"
	"github.com/astaxie/beego"
)

type Authcontroller struct {
	beego.Controller
}
func (c *Authcontroller) Login()  {
	if user := c.GetSession("User"); user != nil {
		c.Data["cuser"] = user
		c.Redirect("/",302)
		return
	}
	if c.Ctx.Input.IsGet(){
		c.TplName = "login.html"
	}else {
		username := c.Input().Get("username")
		passwd := c.Input().Get("password")
		if Udb.Auth(username,passwd) {
			c.SetSession("User",username)
			c.Data["cuser"] = username
			c.Redirect("/",302)
		}else {
			c.Data["form"] = username
			c.Data["errors"] = "用户名或密码错误"
			c.TplName = "login.html"
		}
	}
}

func (c *Authcontroller) Logout() {
	c.DestroySession()
	c.Redirect("/authcontroller/login/", 302)
}

func (c *Authcontroller) Islogin()  {
	if user := c.GetSession("User"); user != nil {
		fmt.Println(user)
		c.Data["cuser"] = user
		return
	}else {
		c.Redirect("/authcontroller/login/", 302)
		return
	}
}