package routers

import (
	"userMgr/controllers"

	beego "github.com/astaxie/beego"
)

func init() {

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.Router("/", &controllers.MainController{})
	beego.Router("/suo/", &controllers.SuoController{})
	beego.Router("/add/:id", &controllers.SuoController{})
	beego.SetStaticPath("/down", "download")
}
