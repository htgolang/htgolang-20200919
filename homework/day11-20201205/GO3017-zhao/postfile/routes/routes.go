package routes

import (
	"net/http"
	"zhao/handles"
)

func init() {
	http.HandleFunc("/file", handles.Postfile)
	http.HandleFunc("/", handles.List)
	http.HandleFunc("/result", handles.Result)

}
