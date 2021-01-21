package routers

import (
	"userMgr/controllers"

	beego "github.com/astaxie/beego"
)

func init() {
	beego.BConfig.WebConfig.DirectoryIndex = true

	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.UserController{})
}
