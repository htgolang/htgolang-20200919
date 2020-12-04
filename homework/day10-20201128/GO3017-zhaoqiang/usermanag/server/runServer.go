package server

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"sync"
	"text/template"
	"time"
	"zhao/models"
	"zhao/utils"
)

func HttpServer(webroot string) {
	local := new(sync.Mutex)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles(filepath.Join(webroot, "userlist.html"))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		users, err := models.File.Load()
		if err != nil {
			fmt.Println("==[models.File]==", err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}

		err = t.ExecuteTemplate(w, "userlist.html", users)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		addr := r.FormValue("addr")
		tel := r.FormValue("tel")
		brithday := r.FormValue("brithday")
		passwd := r.FormValue("passwd")
		e := r.FormValue("err")
		switch {
		case r.Method == "GET":
			t, err := template.ParseFiles(filepath.Join(webroot, "add.html"))
			if err != nil {
				fmt.Println("[template.ParseFile]", err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			u := &models.User{
				Name:     name,
				Addr:     addr,
				Tel:      tel,
				Brithday: brithday,
				Passwd:   passwd,
			}
			a := struct {
				*models.User
				Err string
			}{u, e}
			err = t.ExecuteTemplate(w, "add.html", a)
			if err != nil {
				fmt.Println("[template.ExecuteTemplate]", err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
		case r.Method == "POST":
			users, err := models.File.Load()
			if err != nil {
				fmt.Println("[/add]", err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			var id int
			local.Lock()
			for _, user := range users {
				if user.ID > id {
					id = user.ID
				}
			}
			id++
			local.Unlock()
			r.ParseForm()

			_, err = time.Parse("2006/01/02", brithday)
			if err != nil {
				fmt.Println("time.Parse", err)
				http.Redirect(w, r, fmt.Sprintf("/add?name=%s&passwd=%s&tel=%s&addr=%s&brithday=%s&err=%s", name, passwd, tel, addr, brithday, err.Error()), 301)
				// w.WriteHeader(601)
				// w.Write([]byte(err.Error()))
				w.Write([]byte("111111111111"))
				return
			}

			if name == "" || tel == "" || passwd == "" || brithday == "" {
				w.WriteHeader(601)
				w.Write([]byte("名字， 电话， 密码， 生日不能为空"))
				return
			}
			user := &models.User{
				ID:       id,
				Name:     name,
				Addr:     addr,
				Tel:      tel,
				Brithday: brithday,
				Passwd:   utils.Md5Convert(passwd),
			}
			users = append(users, user)
			err = models.File.Storage(users)
			if err != nil {
				fmt.Println("/add", err)
				w.WriteHeader(500)
				w.Write([]byte("数据存储错误"))
				return
			}
			http.Redirect(w, r, "/", 301)
		}

	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		users, err := models.File.Load()
		if err != nil {
			fmt.Println("[/delete--models.File]", err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			fmt.Println("[/delete]", err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		tmpusers := make([]*models.User, 0, len(users)-1)
		for _, user := range users {
			if user.ID == id {
				continue
			}
			tmpusers = append(tmpusers, user)
		}
		local.Lock()
		err = models.File.Storage(tmpusers)
		local.Unlock()
		if err != nil {
			fmt.Println("[/delete]", err)
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
	})
	u := &models.User{}
	http.HandleFunc("/modify", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			id, err := strconv.Atoi(r.FormValue("id"))
			if err != nil {
				fmt.Println("[strconv.Atoi]", err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			users, err := models.File.Load()
			if err != nil {
				fmt.Println("[models.File.Load]", err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}

			index := -1

			for i, user := range users {
				if user.ID == id {
					index = i
				}
			}
			if index == -1 {
				w.WriteHeader(500)
				w.Write([]byte("用户ID不存在"))
				return
			}

			t, err := template.ParseFiles(filepath.Join(webroot, "modify.html"))
			if err != nil {
				fmt.Println("[template.ParseFile]", err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			ut := struct {
				*models.User
				error
			}{users[index], errors.New(r.FormValue("err"))}
			fmt.Println("=====", index, r.Method, r.Form, r.FormValue("id"))
			err = t.ExecuteTemplate(w, "modify.html", ut)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			u = users[index]
		}
		if r.Method == "POST" {
			name := r.FormValue("name")
			passwd := r.FormValue("passwd")
			tel := r.FormValue("tel")
			addr := r.FormValue("addr")
			brithday := r.FormValue("brithday")
			fmt.Println("++", brithday)
			_, err := time.Parse("2006/01/02", brithday)
			if err != nil {
				fmt.Println("time.Parse", err)
				http.Redirect(w, r, fmt.Sprintf("/modify?id=%d&err=%v", u.ID, err), 301)
				return
			}

			if passwd != u.Passwd {
				passwd = utils.Md5Convert(passwd)
			}

			mdifyuser := &models.User{
				ID:       u.ID,
				Name:     name,
				Addr:     addr,
				Tel:      tel,
				Brithday: brithday,
				Passwd:   passwd,
			}
			users, err := models.File.Load()
			if err != nil {
				fmt.Println("[models.File.Load]", err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			for i, user := range users {
				if user.ID == u.ID {
					users[i] = mdifyuser
					continue
				}
				if name == u.Name {
					http.Redirect(w, r, fmt.Sprintf("/modify?id=%d&err=%v", u.ID, "name重复"), 302)
					return
				}
			}
			err = models.File.Storage(users)
			if err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			http.Redirect(w, r, "/", 301)
		}
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("==[httpServer]==", err)
		return
	}
}
