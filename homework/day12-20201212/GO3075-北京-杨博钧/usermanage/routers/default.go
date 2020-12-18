package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"usermanage/controllers"
)

func Register() {
	beego.AutoRouter(&controllers.UserController{})
	beego.Any("/", func(context *context.Context) {
		context.Redirect(302, "/user/mainpage")
	})
}