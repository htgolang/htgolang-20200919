package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"usermanager_webdb/users"
)

func Queryuser(w http.ResponseWriter, r *http.Request) {
	wlog(r)
	type mydata struct{
		Ud users.Userinfo
		Ok bool
	}
	if r.Method == "POST" {
		wlog(r)
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
		}
		str := r.PostForm.Get("querystr")
		u,ok := Udb.QueryUser(str)
		tpl,err := template.ParseFiles(Templatedir+"/query.html")
		if err != nil {
			fmt.Println(err)
		}
		rdata := mydata{
			Ud: u,
			Ok: ok,
		}
		if err = tpl.ExecuteTemplate(w,"query.html",rdata ); err != nil {
			fmt.Println(err)
		}
	}else {
		wlog(r)
		tpl,err := template.ParseFiles(Templatedir+"/query.html")
		if err != nil {
			fmt.Println(err)
		}
		rdata := mydata{
			Ud: users.Userinfo{},
			Ok: false,
		}
		if err = tpl.ExecuteTemplate(w,"query.html",rdata ); err != nil {
			fmt.Println(err)
		}
	}
}
