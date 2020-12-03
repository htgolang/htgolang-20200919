package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
		}
		un := r.PostForm.Get("username")
		ua := r.PostForm.Get("addr")
		ut := r.PostForm.Get("tel")
		up := r.PostForm.Get("passwd")
		ub := r.PostForm.Get("brth")
		if err := Udb.Add(un, ua, ut, up, ub); err != nil {
			rs := strings.NewReader(fmt.Sprintf("%s",err))
			_,_ = io.Copy(w,rs)
			return
		} else {
			if err = Udb.Sync(); err == nil {
				http.Redirect(w, r, "/", 302)
			}
	
		}

	}else {
		f,err := os.Open(Templatedir+"/add.html")
		if err != nil {
			fmt.Println(err)
		}

		if _, err := io.Copy(w,f); err != nil {
			fmt.Println(err)
		}

	}

}
