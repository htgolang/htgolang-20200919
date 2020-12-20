package routers

import (
	"CMS/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{}, "*:Query")
	beego.AutoRouter(&controllers.UserController{})
}
