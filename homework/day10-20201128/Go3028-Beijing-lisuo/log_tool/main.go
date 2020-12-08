package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.New("home").ParseFiles("./home.html")
		if err != nil {
			panic(err)
		}
		if err := tpl.ExecuteTemplate(w, "home.html", ""); err != nil {
			panic(err)
		}

	})

	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024 * 1024)
		f, fr, err := r.FormFile("filename")
		fmt.Println("f: ", f)
		fmt.Println("fr: ", fr)
		fmt.Println("err: ", err)

	})

	if err := http.ListenAndServe(":8889", nil); err != nil {
		panic(err)
	}

}
