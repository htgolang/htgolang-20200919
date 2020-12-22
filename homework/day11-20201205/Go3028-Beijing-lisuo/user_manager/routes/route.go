package routes

import (
	"net/http"
	controller "user_manager/controllers"
)

func Route() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/create/", controller.CreateUser)
	//http.HandleFunc("/edit/", controller.EditUser)
	//http.HandleFunc("/delete/", controller.DeleteUser)
	//http.HandleFunc("/query/", controller.QueryUser)
}
