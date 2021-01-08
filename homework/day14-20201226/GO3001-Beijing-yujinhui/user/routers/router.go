package routers

import (
	"github.com/astaxie/beego"
	"user/controllers"
)

func init() {
	beego.AutoRouter(&controllers.AuthController{})
	beego.Router("/", &controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})
}
