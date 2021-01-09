package routers

import (
	"user/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// 404
	// beego.ErrorHandler("404", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "404")
	// })
	// beego.ErrorHandler("500", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "500")
	// })
	beego.ErrorController(&controllers.ErrorController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.UserController{})
}
