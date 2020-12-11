package routers

import (
	"CMS/controllers"
	"net/http"
)

func Register() {
	http.HandleFunc("/", controllers.GetUsers)
	http.HandleFunc("/create/", controllers.AddUser)
	http.HandleFunc("/delete/", controllers.DeleteUser)
	http.HandleFunc("/modify/", controllers.ModifyUser)
}
