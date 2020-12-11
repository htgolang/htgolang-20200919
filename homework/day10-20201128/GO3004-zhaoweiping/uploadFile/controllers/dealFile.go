package controllers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"uploadFile/utils"
)

var path = "/Users/aomine/Documents/projects/scripts/go/htgolang-20200919/homework/day10-20201128/GO3004-zhaoweiping/uploadFile/files"

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.RemoteAddr)
	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("template/upload.html"))
		tpl.ExecuteTemplate(w, "upload.html", nil)
	} else if r.Method == "POST" {
		file, handle, err := r.FormFile("file")
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.OpenFile("./files/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(f, file)
		defer f.Close()
		defer file.Close()
		fmt.Println("upload success!")
		http.Redirect(w, r, "/", 302)

	}
}

func ListFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("template/file.html"))
		tpl.ExecuteTemplate(w, "file.html", sortFile())
	} else if r.Method == "POST" {
		http.Redirect(w, r, "/", 302)
	}
}

func sortFile() utils.ByModTime {

	files := utils.SortFile(path)
	return files
}

func StaticFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
}

func StatusFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
}
