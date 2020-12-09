package controller

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ModifyUser(w http.ResponseWriter, r *http.Request)  {
	wlog(r)
	vars := r.URL.Query()
	id := vars.Get("Id")
	if id == "" {
		rs := strings.NewReader("Id 不能为空")
		_,_ = io.Copy(w,rs)
		return
	}
	uid, _ := strconv.Atoi(id)
	uinfo, err := Udb.FindByid(uid)
	if err != nil {
		rs := strings.NewReader(fmt.Sprintf("%s",err))
		_,_ = io.Copy(w,rs)
		return
	}
	if r.Method == "GET" {
		wlog(r)
		tpl,err := template.ParseFiles(Templatedir+"/modify.html")
		if err != nil {
			fmt.Println(err)
		}
		if err = tpl.ExecuteTemplate(w,"modify.html",uinfo); err != nil {
			fmt.Println(err)
		}
	}else {
		wlog(r)
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
		}
		vars := r.URL.Query()
		id := vars.Get("Id")
		name := r.PostForm.Get("username")
		addr := r.PostForm.Get("addr")
		tel := r.PostForm.Get("tel")
		sex:= r.PostForm.Get("sex")
		up := r.PostForm.Get("password")
		ub := r.PostForm.Get("brth")
		err := Udb.Modify(id,name,addr,sex,tel,ub,up)
		if err != nil {
			rs := strings.NewReader(fmt.Sprintf("%s",err))
			_,_ = io.Copy(w,rs)
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}
