package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"text/template"
)

type Result struct {
	Rank   int
	IP     string
	Status int
	URL    string
}

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

	http.HandleFunc("/uploadPage/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.New("upload").ParseFiles("./upload.html")
		if err != nil {
			panic(err)
		}
		if err := tpl.ExecuteTemplate(w, "upload.html", ""); err != nil {
			panic(err)
		}

	})

	http.HandleFunc("/upload/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.URL: ", r.URL)
		// file size limited to 1G
		r.ParseMultipartForm(1000 << 20)
		fmt.Println("r.Form: ", r.Form)
		fmt.Println("r.Postform: ", r.PostForm)
		fmt.Println("r.MultipartForm: ", r.MultipartForm)
		if fileHeaders, ok := r.MultipartForm.File[r.Form["filename"][0]]; ok {
			for _, fileHeader := range fileHeaders {
				fmt.Println("fileHeader.Filename: ", fileHeader.Filename)
				fmt.Println("fileHeader.Size: ", fileHeader.Size)
				newFile, err := os.Create("/tmp/" + fileHeader.Filename)
				if err != nil {
					panic(err)
				}
				file, err := fileHeader.Open()
				if err != nil {
					panic(err)
				}
				io.Copy(newFile, file)
			}
		}
		http.Redirect(w, r, "/", 302)

	})

	http.HandleFunc("/resultPage/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("resultPage==========")
		tpl, err := template.New("result").ParseFiles("./result.html")
		if err != nil {
			panic(err)
		}
		if err := tpl.ExecuteTemplate(w, "result.html", ""); err != nil {
			panic(err)
		}
		fmt.Println("resultPage==========end")
	})

	http.HandleFunc("/display/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.New("display").ParseFiles("./display.html")
		if err != nil {
			panic(err)
		}
		resultList := []Result{
			Result{1, "192.178.20.30", 500, "/"},
			Result{2, "172.12.50.35", 200, "/cake/bake"},
		}
		if err := tpl.ExecuteTemplate(w, "display.html", resultList); err != nil {
			panic(err)
		}

	})

	if err := http.ListenAndServe(":8889", nil); err != nil {
		panic(err)
	}

}
