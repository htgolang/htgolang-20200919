package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/cmd/funcs"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/define"
)

var (
	addr = ":8889"
	abs  = "/data/htgolang-20200919/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/web"
)

type User struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Cell    string    `json:"cell"`
	Born    time.Time `json:"born"`
	Passwd  string    `json:"passwd"`
}

func Serv() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("current exec dir: ", dir)
	users := &define.UserList
	fmt.Println("users: ", *users)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("home").ParseFiles(abs + "/html/home.html")
		if err != nil {
			panic(err)
		}
		t.ExecuteTemplate(w, "home.html", users)
	})

	http.HandleFunc("/create/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, err := template.New("create").ParseFiles(abs + "/html/create.html")
			if err != nil {
				panic(err)
			}
			t.ExecuteTemplate(w, "create.html", nil)
		} else {
			r.ParseForm()
			fmt.Println("r.PostForm from /create/:", r.PostForm)
			i, err := strconv.ParseInt(r.PostForm["id"][0], 10, 64)
			if err != nil {
				panic(err)
			}
			// error
			if i == 0 {
				t, err := template.New("error").ParseFiles(abs + "/html/error.html")
				if err != nil {
					panic(err)
				}
				t.ExecuteTemplate(w, "error.html", "id 0 belongs to admin......")
				http.Redirect(w, r, "/", 302)
				return
			}
			t, errP := time.Parse("2006.01.02", r.PostForm["born"][0])
			if errP != nil {
				panic(errP)
			}
			userN := define.User{
				ID:      i,
				Name:    r.PostForm["name"][0],
				Address: r.PostForm["address"][0],
				Cell:    r.PostForm["cell"][0],
				Born:    t,
				Passwd:  r.PostForm["passwd"][0],
			}
			*users = append(*users, userN)
			// error
			if r.PostForm["name"][0] == "jack" {
				t, err := template.New("error").ParseFiles(abs + "/html/error.html")
				if err != nil {
					panic(err)
				}
				t.ExecuteTemplate(w, "error.html", "do not create jack......")
			}
			fmt.Println("userToCreate:", userN)
			http.Redirect(w, r, "/", 302)
		}

	})

	http.HandleFunc("/edit/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, err := template.New("edit").ParseFiles(abs + "/html/edit.html")
			if err != nil {
				panic(err)
			}
			r.ParseForm()
			fmt.Println("r.Form from /edit/:", r.Form)
			i := r.Form["id"][0]
			id, err := strconv.ParseInt(i, 10, 64)
			if err != nil {
				panic(err)
			}
			user, err := funcs.IDFindUser(users, id)
			fmt.Println("use to mod: ", user)
			if err != nil {
				t.ExecuteTemplate(w, "error.html", "no such user......")
				panic(err)
			}
			// error
			if r.Form["id"][0] == "0" {
				t, err := template.New("error").ParseFiles(abs + "/html/error.html")
				if err != nil {
					panic(err)
				}
				t.ExecuteTemplate(w, "error.html", "do not edit admin......")
				return
			}
			t.ExecuteTemplate(w, "edit.html", user)
		} else if r.Method == "POST" {
			var index int
			var id int64
			r.ParseForm()
			fmt.Println("from /edit/ r.PostFrom: ", r.PostForm)
			id, err = strconv.ParseInt(r.PostForm["id"][0], 10, 64)
			if err != nil {
				panic(err)
			}
			fmt.Println("id: ", id)
			for i, user := range *users {
				if user.ID == id {
					index = i
					userO := (*users)[index]
					userN := define.User{
						ID:      userO.ID,
						Name:    r.PostForm["name"][0],
						Address: r.PostForm["address"][0],
						Cell:    r.PostForm["cell"][0],
						Born:    userO.Born,
						Passwd:  r.PostForm["passwd"][0],
					}
					(*users)[index] = userN
				}
			}
			http.Redirect(w, r, "/", 302)
		}

	})

	http.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("delete").ParseFiles(abs + "/html/delete.html")
		if err != nil {
			panic(err)
		}
		r.ParseForm()
		fmt.Println("r.Form from /delete/:", r.Form)
		i, errP := strconv.ParseInt(r.Form["id"][0], 10, 64)
		if errP != nil {
			panic(errP)
		}
		u := func(users *[]define.User) define.User {
			for _, user := range *users {
				if user.ID == i {
					return user
				}
			}
			return (*users)[0]
		}
		// error
		if r.Form["id"][0] == "0" {
			t, err := template.New("error").ParseFiles(abs + "/html/error.html")
			if err != nil {
				panic(err)
			}
			t.ExecuteTemplate(w, "error.html", "do not delete admin......")
			http.Redirect(w, r, "/", 302)
			return
		}
		fmt.Println("userToDel: ", u(users))
		t.ExecuteTemplate(w, "delete.html", u(users))

	})

	http.HandleFunc("/query/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			t, err := template.New("query").ParseFiles(abs + "/html/query.html")
			if err != nil {
				panic(err)
			}
			t.ExecuteTemplate(w, "query.html", nil)
		} else if r.Method == "POST" {
			r.ParseForm()
			fmt.Println("r.PostForm from /edit/: ", r.PostForm)
			var gotUsers []define.User
			inputList := []string{
				r.PostForm["id"][0],
				r.PostForm["name"][0],
				r.PostForm["address"][0],
				r.PostForm["born"][0],
			}
			for _, user := range *users {
				for _, input := range inputList {
					b := func(u define.User, input string) bool {
						return strings.Contains(strings.ToLower(u.Name), input) ||
							strings.Contains(strings.ToLower(u.Address), input) ||
							strings.Contains(u.Cell, input) ||
							strings.Contains(u.Born.Format("2006.01.02"), input)
					}(user, input)
					if b {
						gotUsers = append(gotUsers, user)
					}
				}
			}
			// display
			t, err := template.New("display").ParseFiles(abs + "/html/display.html")
			if err != nil {
				panic(err)
			}
			t.ExecuteTemplate(w, "display.html", gotUsers)
		}
	})

	http.ListenAndServe(addr, nil)
	fmt.Println("Server start up at: ", addr)
}
