package controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		wlog(r)
		if err := r.ParseForm(); err != nil {
			log.Fatalln(err)
		}
		un := r.PostForm.Get("username")
		ua := r.PostForm.Get("addr")
		us := r.PostForm.Get("sex")
		ut := r.PostForm.Get("tel")
		up := r.PostForm.Get("password")
		ub := r.PostForm.Get("brth")
		if err := Udb.Add(un, ua,us, ut, up, ub); err != nil {
			rs := strings.NewReader(fmt.Sprintf("%s",err))
			_,_ = io.Copy(w,rs)
			return
		} else {
			http.Redirect(w, r, "/", 302)
		}

	}else {
		wlog(r)
		f,err := os.Open(Templatedir+"/add.html")
		if err != nil {
			fmt.Println(err)
		}

		if _, err := io.Copy(w,f); err != nil {
			fmt.Println(err)
		}

	}

}
