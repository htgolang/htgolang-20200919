package controllers

import (
	"net/http"
	"strconv"
	"text/template"
	"user/services"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("template/user.html"))
	tpl.ExecuteTemplate(w, "user.html", services.GetUsers())
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl := template.Must(template.ParseFiles("template/create.html"))
		tpl.ExecuteTemplate(w, "create.html", nil)
	} else {
		services.AddUser(
			r.FormValue("name"),
			r.FormValue("addr"),
			r.FormValue("sex") == "1",
		)
		http.Redirect(w, r, "/", 302)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if id, err := strconv.ParseInt(r.FormValue("id"), 10, 64); err == nil {
		services.DeleteUser(id)
	}
	http.Redirect(w, r, "/", 302)
}
