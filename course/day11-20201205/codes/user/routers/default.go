package routers

import (
	"net/http"
	"user/controllers"
)

func Register() {
	http.HandleFunc("/", controllers.GetUsers)
	http.HandleFunc("/create/", controllers.AddUser)
	http.HandleFunc("/delete/", controllers.DeleteUser)
}
