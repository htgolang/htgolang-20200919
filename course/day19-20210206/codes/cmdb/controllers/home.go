package controllers

import (
	"net/http"

	"github.com/astaxie/beego"

	base "cmdb/base/controllers"
)

type HomeController struct {
	base.AuthorizationController
}

func (c *HomeController) Index() {
	c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}
