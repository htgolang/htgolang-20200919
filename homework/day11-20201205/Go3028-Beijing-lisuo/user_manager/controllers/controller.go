package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"
	"user_manager/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("home").ParseFiles("template/home.html")
	if err != nil {
		panic(err)
	}
	users, err := services.ListAllUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	t.ExecuteTemplate(w, "home.html", users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("create").ParseFiles("template/create.html")
		if err != nil {
			panic(err)
		}
		t.ExecuteTemplate(w, "create.html", nil)
	} else {
		r.ParseForm()
		sexInt, err := strconv.Atoi(r.PostForm["sex"][0])
		t, errP := time.Parse("2006.01.02", r.PostForm["born"][0])
		if err != nil {
			if errP != nil {
				ErrorMsg(w, "Sex bad format, 1 for male 0 for female\nBorn bad format, born time example: 1995.07.07")
				return
			}
			return
		} else if errP != nil {
			ErrorMsg(w, "Born bad format, born time example: 1995.07.07")
			return
		}
		var (
			Name    = r.PostForm["name"][0]
			Sex     = sexInt
			Address = r.PostForm["address"][0]
			Cell    = r.PostForm["cell"][0]
			Born    = t
			Passwd  = r.PostForm["passwd"][0]
		)
		if r.PostForm["name"][0] == "admin" {
			ErrorMsg(w, "Do not create admin.")
			return
		}
		fmt.Println("userToCreate:", Name)
		if err := services.CreateUser(Name, Passwd, Address, Cell, Sex, Born); err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}

//
//func EditUser(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "GET" {
//		t, err := template.New("edit").ParseFiles(abs + "/html/edit.html")
//		if err != nil {
//			panic(err)
//		}
//		r.ParseForm()
//		fmt.Println("r.Form from /edit/:", r.Form)
//		i := r.Form["id"][0]
//		id, err := strconv.ParseInt(i, 10, 64)
//		if err != nil {
//			panic(err)
//		}
//		user, err := funcs.IDFindUser(users, id)
//		fmt.Println("use to mod: ", user)
//		if err != nil {
//			t.ExecuteTemplate(w, "error.html", "no such user......")
//			panic(err)
//		}
//		// error
//		if r.Form["id"][0] == "0" {
//			t, err := template.New("error").ParseFiles(abs + "/html/error.html")
//			if err != nil {
//				panic(err)
//			}
//			t.ExecuteTemplate(w, "error.html", "do not edit admin......")
//			return
//		}
//		t.ExecuteTemplate(w, "edit.html", user)
//	} else if r.Method == "POST" {
//		var index int
//		var id int64
//		r.ParseForm()
//		fmt.Println("from /edit/ r.PostFrom: ", r.PostForm)
//		id, err = strconv.ParseInt(r.PostForm["id"][0], 10, 64)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println("id: ", id)
//		for i, user := range *users {
//			if user.ID == id {
//				index = i
//				userO := (*users)[index]
//				userN := define.User{
//					ID:      userO.ID,
//					Name:    r.PostForm["name"][0],
//					Address: r.PostForm["address"][0],
//					Cell:    r.PostForm["cell"][0],
//					Born:    userO.Born,
//					Passwd:  r.PostForm["passwd"][0],
//				}
//				(*users)[index] = userN
//			}
//		}
//		http.Redirect(w, r, "/", 302)
//	}
//}
//
//func DeleteUser(w http.ResponseWriter, r *http.Request) {
//	t, err := template.New("delete").ParseFiles(abs + "/html/delete.html")
//	if err != nil {
//		panic(err)
//	}
//	r.ParseForm()
//	fmt.Println("r.Form from /delete/:", r.Form)
//	i, errP := strconv.ParseInt(r.Form["id"][0], 10, 64)
//	if errP != nil {
//		panic(errP)
//	}
//	u := func(users *[]define.User) define.User {
//		for _, user := range *users {
//			if user.ID == i {
//				return user
//			}
//		}
//		return (*users)[0]
//	}
//	// error
//	if r.Form["id"][0] == "0" {
//		t, err := template.New("error").ParseFiles(abs + "/html/error.html")
//		if err != nil {
//			panic(err)
//		}
//		t.ExecuteTemplate(w, "error.html", "do not delete admin......")
//		http.Redirect(w, r, "/", 302)
//		return
//	}
//	fmt.Println("userToDel: ", u(users))
//	t.ExecuteTemplate(w, "delete.html", u(users))
//}
//
//func QueryUser(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "GET" {
//		t, err := template.New("query").ParseFiles(abs + "/html/query.html")
//		if err != nil {
//			panic(err)
//		}
//		t.ExecuteTemplate(w, "query.html", nil)
//	} else if r.Method == "POST" {
//		r.ParseForm()
//		fmt.Println("r.PostForm from /edit/: ", r.PostForm)
//		var gotUsers []define.User
//		inputList := []string{
//			r.PostForm["id"][0],
//			r.PostForm["name"][0],
//			r.PostForm["address"][0],
//			r.PostForm["born"][0],
//		}
//		for _, user := range *users {
//			for _, input := range inputList {
//				b := func(u define.User, input string) bool {
//					return strings.Contains(strings.ToLower(u.Name), input) ||
//						strings.Contains(strings.ToLower(u.Address), input) ||
//						strings.Contains(u.Cell, input) ||
//						strings.Contains(u.Born.Format("2006.01.02"), input)
//				}(user, input)
//				if b {
//					gotUsers = append(gotUsers, user)
//				}
//			}
//		}
//		// display
//		t, err := template.New("display").ParseFiles(abs + "/html/display.html")
//		if err != nil {
//			panic(err)
//		}
//		t.ExecuteTemplate(w, "display.html", gotUsers)
//	}
//}

func ErrorMsg(w http.ResponseWriter, msg string) {
	t, err := template.New("error").ParseFiles("template/error.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "error.html", msg)
}
