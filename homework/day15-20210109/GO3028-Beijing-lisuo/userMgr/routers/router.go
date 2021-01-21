package routers

import (
	"userMgr/controllers"

	beego "github.com/astaxie/beego"
)

func init() {
	beego.BConfig.WebConfig.DirectoryIndex = true

	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.AuthController{})
}
