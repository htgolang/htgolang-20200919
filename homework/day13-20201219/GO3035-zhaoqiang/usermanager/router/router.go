package router

import (
	"user/controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", new(controller.UserController), "*:Jump")
	beego.AutoRouter(new(controller.UserController))
	beego.AutoRouter(new(controller.AuthController))
	// beego.AutoRouter(new(controller.BaseController))
	beego.AutoRouter(new(controller.PermissionController))
}
