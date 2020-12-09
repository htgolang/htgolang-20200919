package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"usermanager_webdb/users"
)

var Udb users.Mydb
var Templatedir string

func Index(w http.ResponseWriter, r *http.Request) {
	wlog(r)
	tpl,err := template.ParseFiles(Templatedir+"/index.html")
	if err != nil {
		fmt.Println(err)
	}
	if err = tpl.ExecuteTemplate(w,"index.html",Udb.Getall()); err != nil {
		fmt.Println(err)
	}

}

