package routers

import (
	"net/http"
	"uploadFile/controllers"
)

func Register() {
	http.HandleFunc("/", controllers.ListFile)
	http.HandleFunc("/upload", controllers.UploadFile)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./files"))))
	http.HandleFunc("/status", controllers.StatusFile)
}
