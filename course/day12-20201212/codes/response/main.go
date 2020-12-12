package main

import (
	"github.com/astaxie/beego"
)

/*
Ctx
	c.Ctx
	c.Ctx.ResponseWirter
	c.Ctx.Output
Controller
*/

type ResponseController struct {
	beego.Controller
}

func (c *ResponseController) Test() {
	c.Data["user"] = "kk"
	c.TplName = "responsecontroller/test.html"
}

func main() {
	beego.AutoRouter(&ResponseController{})
	beego.Run()
}
