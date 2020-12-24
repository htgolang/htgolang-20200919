package route

import (
	"net/http"
	"webtool/controllers"
)

func Route() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/upload/", controllers.Upload)
	http.HandleFunc("/uploadPage/", controllers.UploadPage)
	http.HandleFunc("/resultPage/", controllers.ResultPage)
	http.HandleFunc("/display/", controllers.Display)
}
