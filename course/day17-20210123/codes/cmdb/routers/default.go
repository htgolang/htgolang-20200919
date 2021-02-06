package routers

import (
	"cmdb/controllers"
	"cmdb/controllers/api"
	"cmdb/filters"

	"github.com/astaxie/beego"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	// 404
	// beego.ErrorHandler("404", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "404")
	// })
	// beego.ErrorHandler("500", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "500")
	// })

	beego.InsertFilter("/*", beego.BeforeExec, filters.BeferExec)
	beego.InsertFilter("/*", beego.AfterExec, filters.AfterExec, false)

	beego.ErrorController(&controllers.ErrorController{})

	beego.Handler("/metrics/", promhttp.Handler())

	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.UserController{})

	beego.AutoRouter(&api.PrometheusController{})

}
