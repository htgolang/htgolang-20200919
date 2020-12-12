package routers

import (
	"net/http"
	"user/controllers"
	"user/logger"
)

func Register() {
	http.HandleFunc("/", logger.LoggerWrapper(controllers.GetUsers))
	http.HandleFunc("/create/", logger.LoggerWrapper(controllers.AddUser))
	http.HandleFunc("/delete/", logger.LoggerWrapper(controllers.DeleteUser))
}
