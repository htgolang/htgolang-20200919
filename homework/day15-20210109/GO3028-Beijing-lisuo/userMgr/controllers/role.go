package controllers

import (
	"userMgr/services"

	beego "github.com/astaxie/beego"
)

type RoleController struct {
	beego.Controller
}

func (c *RoleController) Role() {
	if services.Admin() {
		c.TplName = "role/role.html"
	}

}
