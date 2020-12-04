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
	vars := r.URL.Query()
	id := vars.Get("Id")
	if id == "" {
		rs := strings.NewReader("Id 不能为空")
		_,_ = io.Copy(w,rs)
		return
	}
	uid, _ := strconv.Atoi(id)
	uinfo, index, err := Udb.FindByid(uid)
	if err != nil {
		rs := strings.NewReader(fmt.Sprintf("%s",err))
		_,_ = io.Copy(w,rs)
		return
	}
	if r.Method == "GET" {
		tpl,err := template.ParseFiles(Templatedir+"/modify.html")
		if err != nil {
			fmt.Println(err)
		}
		if err = tpl.ExecuteTemplate(w,"modify.html",uinfo); err != nil {
			fmt.Println(err)
		}
	}else {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
		}
		name := r.PostForm.Get("username")
		addr := r.PostForm.Get("addr")
		tel := r.PostForm.Get("tel")
		up := r.PostForm.Get("passwd")
		ub := r.PostForm.Get("brth")
		err := Udb.Modify(index,name,addr,tel,ub,up)
		if err != nil {
			rs := strings.NewReader(fmt.Sprintf("%s",err))
			_,_ = io.Copy(w,rs)
			return
		}
		if err = Udb.Sync(); err != nil {
			rs := strings.NewReader(fmt.Sprintf("%s",err))
			_,_ = io.Copy(w,rs)
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}
