package routers

import (
	"net/http"
	"zhao/controllers"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	http.HandleFunc("/", controllers.WrapHandler(controllers.GetUsers))
	http.HandleFunc("/create", controllers.WrapHandler(controllers.AddUser))
	http.HandleFunc("/delete", controllers.WrapHandler(controllers.DeleteUser))
	http.HandleFunc("/modify", controllers.WrapHandler(controllers.ModifyUser))
	http.Handle("/metrics", promhttp.Handler())
}
