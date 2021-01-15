package controller

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404()  {
	c.Data["content"] = "Page not found"
	c.TplName = "error/404.html"
}

func (c *ErrorController) Error403()  {
	c.Data["content"] = "No Permission"
	c.TplName = "error/404.html"
}