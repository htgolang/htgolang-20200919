package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
	"user_manager/models"
	"user_manager/services"
	"user_manager/user_utils"
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
		} else if err := services.NameFindUser(r.PostFormValue("name")); err == nil {
			ErrorMsg(w, "Error: There's a user with same name: "+r.PostFormValue("name"))
			return
		}
		if !user_utils.JustDigits(r.PostFormValue("cell")) {
			ErrorMsg(w, "Error: Cell must be pure digits.")
			return
		}
		fmt.Println("userToCreate:", Name)
		if err := services.CreateUser(Name, Passwd, Address, Cell, Sex, Born); err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/", 302)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		ErrorMsg(w, err.Error())
		return
	}
	if id == 5 {
		ErrorMsg(w, "You can't delete admin, who's id is "+r.FormValue("id"))
		return
	}
	services.IDDelUser(id)
	ErrorMsg(w, "deleted user by id: "+r.FormValue("id"))
	http.Redirect(w, r, "/", 302)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	type cUser struct {
		ID      int64
		Name    string
		Sex     int
		Address string
		Cell    string
		Born    string
		Passwd  string
	}
	if r.Method == "GET" {
		t, err := template.New("edit").ParseFiles("template/edit.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		r.ParseForm()
		fmt.Println("r.Form from /edit/:", r.Form)
		i := r.FormValue("id")
		id, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			ErrorMsg(w, err.Error())
			return
		}
		user, errf := services.IDFindUser(id)
		var cuser = cUser{
			ID:      user.ID,
			Name:    user.Name,
			Sex:     user.Sex,
			Address: user.Address,
			Cell:    user.Cell,
			Born:    strings.Split(user.Born.String(), " ")[0],
			Passwd:  user.Passwd,
		}
		if errf != nil {
			ErrorMsg(w, "No such user.")
		}
		t.ExecuteTemplate(w, "edit.html", cuser)
	} else if r.Method == "POST" {
		var id int64
		r.ParseForm()
		fmt.Println("from /edit/ r.PostFrom: ", r.PostForm)
		id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
		if err != nil {
			ErrorMsg(w, err.Error())
			return
		}
		if id == models.AdminID {
			ErrorMsg(w, "Do not edit admin.")
			return
		}
		fmt.Println("edit user, id is: ", id)
		var (
			name     = r.PostFormValue("name")
			address  = r.PostFormValue("address")
			password = r.PostFormValue("passwd")
			cell     = r.PostFormValue("cell")
			sex      = r.PostFormValue("sex")
			born     = r.PostFormValue("born")
		)
		if err := services.IDModUser(name, address, password, cell, sex, born, id); err != nil {
			ErrorMsg(w, err.Error())
			return
		}
		http.Redirect(w, r, "/", 302)
	}
}

// QueryUser get user by some sting
func QueryUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("query").ParseFiles("template/query.html")
		if err != nil {
			panic(err)
		}
		t.ExecuteTemplate(w, "query.html", nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println("r.PostForm from /edit/: ", r.PostForm)
		var (
			id      = r.PostFormValue("id")
			name    = r.PostFormValue("name")
			sex     = r.PostFormValue("sex")
			address = r.PostFormValue("address")
			cell    = r.PostFormValue("cell")
		)
		users, err := services.QueryUser(id, name, sex, address, cell)
		if err != nil {
			ErrorMsg(w, err.Error())
			return
		}
		// display
		t, err := template.New("display").ParseFiles("template/display.html")
		if err != nil {
			panic(err)
		}
		t.ExecuteTemplate(w, "display.html", users)
	}
}

func ErrorMsg(w http.ResponseWriter, msg string) {
	t, err := template.New("error").ParseFiles("template/error.html")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "error.html", msg)
}
