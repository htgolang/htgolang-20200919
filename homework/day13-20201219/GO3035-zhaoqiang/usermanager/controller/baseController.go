package controller

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (base *BaseController) Prepare() {
	session := base.GetSession("user")
	if session == nil {
		base.Redirect("/auth/login", 301)
	}
}
