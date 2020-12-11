package main

import (
	"net/http"
	"uploadFile/routers"
)

func main() {
	routers.Register()
	http.ListenAndServe(":8888", nil)
}
