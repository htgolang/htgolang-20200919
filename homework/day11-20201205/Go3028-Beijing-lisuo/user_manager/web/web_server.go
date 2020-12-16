package web

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
	"time"

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

	users := define.UserList
	//users := []*User{
	//	{0, "jaccy", "Beijing", "18811738208", time.Now(), "sldfjasldfjaldhfgasdlkga"},
	//	{1, "jack ma", "HangZhou", "18888888888", time.Now().Add(4 * time.Hour), "sdklfg adlfjkgaldffasldkfh"},
	//	{2, "steve", "US", "18811733333", time.Now(), "dslfajflfjasdfjkasfl;fdfl;aj"},
	//	{3, "spilman", "German", "18811118208", time.Now().Add(2 * time.Second), "sdklfasl;fjaflkkfa;fjaskl;"},
	//	{4, "Jhon", "Venus", "18811722208", time.Now(), "klsdfjal;dfjasldfjsdklfj"},
	//}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.New("home").ParseFiles(abs + "/html/home.html")
		if err != nil {
			panic(err)
		}
		t.ExecuteTemplate(w, "home.html", users)

		r.ParseForm()
		fmt.Println("r.Form from /: ", r.Form)
		var id int64
		if len(r.Form["id"]) != 0 {
			id, err = strconv.ParseInt(r.Form["id"][0], 10, 64)
			if err != nil {
				panic(err)
			}
		}
		nUsers := make([]define.User, 0, len(users))
		for _, user := range users {
			if user.ID == id {
				continue
			}
			nUsers = append(nUsers, user)
		}
		users = nUsers
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
			users = append(users, userN)
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
		fmt.Println("userToDel: ", u(&users))
		t.ExecuteTemplate(w, "delete.html", u(&users))

	})

	http.ListenAndServe(addr, nil)
}
