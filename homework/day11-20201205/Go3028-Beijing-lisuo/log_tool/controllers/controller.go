package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"webtool/models"
	"webtool/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("home").ParseFiles("template/home.html")
	if err != nil {
		panic(err)
	}
	if err := tpl.ExecuteTemplate(w, "home.html", ""); err != nil {
		panic(err)
	}

}

func UploadPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("upload").ParseFiles("template/upload.html")
	if err != nil {
		panic(err)
	}
	if err := tpl.ExecuteTemplate(w, "upload.html", ""); err != nil {
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// services.Upload()
	fmt.Println("upload==========start")
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
			models.UploadFile = models.UploadDir + fileHeader.Filename
			err := services.Upload(fileHeader, models.UploadFile)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			}
		}
	}
	http.Redirect(w, r, "/", 302)
	fmt.Println("upload==========end")

}

func ResultPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("resultPage==========start")
	tpl, err := template.New("result").ParseFiles("template/result.html")
	if err != nil {
		panic(err)
	}
	if err := tpl.ExecuteTemplate(w, "result.html", ""); err != nil {
		panic(err)
	}
	fmt.Println("resultPage==========end")
}

// process the log file and  display result
func Display(w http.ResponseWriter, r *http.Request) {
	fmt.Println("displayPage==========start")
	tpl, err := template.New("display").ParseFiles("template/display.html")
	if err != nil {
		panic(err)
	}
	// calculate the result
	// services.CalcResult()
	resultList, err := services.CalcResult()
	if err != nil {
		fmt.Fprint(w, err.Error())
		http.Redirect(w, r, "/", 302)
		return
	}

	if err := tpl.ExecuteTemplate(w, "display.html", resultList); err != nil {
		panic(err)
	}
	fmt.Println("displayPage==========end")
}
