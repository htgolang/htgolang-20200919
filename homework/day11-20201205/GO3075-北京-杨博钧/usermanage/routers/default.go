package routers

import (
	"net/http"
	"usermanage/controllers"
)

func Register() {
	http.HandleFunc("/", controllers.MainPage)
	http.HandleFunc("/add/", controllers.AddPage)
	http.HandleFunc("/update/", controllers.UpdatePage)
	http.HandleFunc("/delete/", controllers.DeletePage)
}