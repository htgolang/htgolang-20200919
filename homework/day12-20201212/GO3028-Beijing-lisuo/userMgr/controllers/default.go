package controllers

import (
	beego "github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type SuoController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *SuoController) Get() {
	c.Data["Suosuoli"] = "https://www.suosuoli.cn"
	c.Data["Visitors"] = "500"
	c.TplName = "suosuoli.html"
}
