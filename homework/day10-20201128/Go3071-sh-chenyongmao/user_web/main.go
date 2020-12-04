package main

import (
	"net/http"
	"strconv"
	"text/template"
)

type User struct {
	Id   int
	Name string
	Age  string
	Addr string
	Sex  bool
}

func GetId(u []*User) int {
	id := 0
	for _, user := range u {
		if user.Id > id {
			id = user.Id
		}
	}
	return id + 1
}

func GetUser(u []*User, id string) *User {
	if nid, err := strconv.Atoi(id); err == nil {
		for _, user := range u {
			if user.Id == nid {
				return user
			}
		}
	}
	return nil
}

func main() {
	addr := "0.0.0.0:8080"
	users := []*User{
		{1, "koko", "15", "xxxx", true},
		{2, "jojo", "20", "yyyy", false},
		{3, "lolo", "13", "zzzz", true},
		{4, "bobo", "5", "vvvv", true},
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("template/users.html"))
		tpl.ExecuteTemplate(w, "users.html", users)
	})
	http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		if id, err := strconv.Atoi(r.FormValue("id")); err == nil {
			nUsers := make([]*User, 0, len(users))
			for _, user := range users {
				if user.Id != id {
					nUsers = append(nUsers, user)
				}
				users = nUsers
			}
		}
		http.Redirect(w, r, "/", 302)
	})
	http.HandleFunc("/add/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			tpl := template.Must(template.ParseFiles("template/create_user.html"))
			tpl.ExecuteTemplate(w, "create_user.html", nil)
		} else {
			users = append(users, &User{
				GetId(users),
				r.FormValue("name"),
				r.FormValue("age"),
				r.FormValue("addr"),
				r.FormValue("sex") == "1",
			})
			http.Redirect(w, r, "/", 302)
		}
	})
	http.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			user := GetUser(users, r.FormValue("id"))
			tpl := template.Must(template.ParseFiles("template/update_user.html"))
			tpl.ExecuteTemplate(w, "update_user.html", user)
		} else {
			if nid, err := strconv.Atoi(r.FormValue("id")); err == nil {
				for k, u := range users {
					if u.Id == nid {
						users[k] = &User{
							nid,
							r.FormValue("name"),
							r.FormValue("age"),
							r.FormValue("addr"),
							r.FormValue("sex") == "1",
						}
					}
				}
			}
			http.Redirect(w, r, "/", 302)
		}

	})
	http.ListenAndServe(addr, nil)
}
