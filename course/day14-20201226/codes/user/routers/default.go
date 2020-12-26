package routers

import (
	"user/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.UserController{})
}
