package controllers

import (
	"fmt"
	"time"
	"userMgr/models"
	"userMgr/services"

	beego "github.com/astaxie/beego"
)

// UserController is the main controller of userMgr
type UserController struct {
	beego.Controller
}

// Home give a default page, with a list of users
func (c *UserController) Home() {
	users, err := services.ListAllUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Data["users"] = users
	c.TplName = "user/home.html"
}

// Create add a user to user list if request method is post
func (c *UserController) Create() {
	if c.Ctx.Input.IsPost() {
		bornStr := c.GetString("born")
		t, errP := time.Parse("2006.01.02", bornStr)
		if errP != nil {
			c.ErrorMsg(errP.Error())
			return
		}
		sex, err := c.GetInt("sex")
		if err != nil {
			c.ErrorMsg(err.Error())
			return
		}
		name := c.GetString("name")
		address := c.GetString("address")
		cell := c.GetString("cell")
		password := c.GetString("passwd")

		// create user here
		if errc := services.CreateUser(name, password, address, cell, sex, t); errc != nil {
			c.ErrorMsg(errc.Error())
			return
		}

		// redirect to home after create user
		c.Redirect("/user/home/", 301)
	} else {
		c.TplName = "user/create.html"
	}
}

func (c *UserController) Delete() {
	id, err := c.GetInt64("id")
	if err != nil {
		c.ErrorMsg(err.Error())
		return
	}
	if id == models.AdminID {
		c.ErrorMsg("You can't delete admin, who's id is: " + c.GetString("id"))
		return
	}
	fmt.Println("To delete: ", c.GetString("id"))

	c.Confirm("Are you sure to delete user: " + c.GetString("id"))

	if errd := services.IDDelUser(id); errd == nil {
		c.ErrorMsg("Deleted a user, who's id is: " + c.GetString("id"))
		return
		//c.Redirect("/user/home/", 301)
	} else {
		c.ErrorMsg(errd.Error())
		return
	}
}

//func DeleteUser(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm()
//	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
//	if err != nil {
//		ErrorMsg(w, err.Error())
//		return
//	}
//	if id == 5 {
//		ErrorMsg(w, "You can't delete admin, who's id is "+r.FormValue("id"))
//		return
//	}
//	services.IDDelUser(id)
//	ErrorMsg(w, "deleted user by id: "+r.FormValue("id"))
//	http.Redirect(w, r, "/", 302)
//}
//
//func EditUser(w http.ResponseWriter, r *http.Request) {
//	type cUser struct {
//		ID      int64
//		Name    string
//		Sex     int
//		Address string
//		Cell    string
//		Born    string
//		Passwd  string
//	}
//	if r.Method == "GET" {
//		t, err := template.New("edit").ParseFiles("template/edit.html")
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		r.ParseForm()
//		fmt.Println("r.Form from /edit/:", r.Form)
//		i := r.FormValue("id")
//		id, err := strconv.ParseInt(i, 10, 64)
//		if err != nil {
//			ErrorMsg(w, err.Error())
//			return
//		}
//		user, errf := services.IDFindUser(id)
//		var cuser = cUser{
//			ID:      user.ID,
//			Name:    user.Name,
//			Sex:     user.Sex,
//			Address: user.Address,
//			Cell:    user.Cell,
//			Born:    strings.Split(user.Born.String(), " ")[0],
//			Passwd:  user.Passwd,
//		}
//		if errf != nil {
//			ErrorMsg(w, "No such user.")
//		}
//		t.ExecuteTemplate(w, "edit.html", cuser)
//	} else if r.Method == "POST" {
//		var id int64
//		r.ParseForm()
//		fmt.Println("from /edit/ r.PostFrom: ", r.PostForm)
//		id, err := strconv.ParseInt(r.PostFormValue("id"), 10, 64)
//		if err != nil {
//			ErrorMsg(w, err.Error())
//			return
//		}
//		if id == models.AdminID {
//			ErrorMsg(w, "Do not edit admin.")
//			return
//		}
//		fmt.Println("edit user, id is: ", id)
//		var (
//			name     = r.PostFormValue("name")
//			address  = r.PostFormValue("address")
//			password = r.PostFormValue("passwd")
//			cell     = r.PostFormValue("cell")
//			sex      = r.PostFormValue("sex")
//			born     = r.PostFormValue("born")
//		)
//		if err := services.IDModUser(name, address, password, cell, sex, born, id); err != nil {
//			ErrorMsg(w, err.Error())
//			return
//		}
//		http.Redirect(w, r, "/", 302)
//	}
//}
//
//// QueryUser get user by some sting
//func QueryUser(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "GET" {
//		t, err := template.New("query").ParseFiles("template/query.html")
//		if err != nil {
//			panic(err)
//		}
//		t.ExecuteTemplate(w, "query.html", nil)
//	} else if r.Method == "POST" {
//		r.ParseForm()
//		fmt.Println("r.PostForm from /edit/: ", r.PostForm)
//		var (
//			id      = r.PostFormValue("id")
//			name    = r.PostFormValue("name")
//			sex     = r.PostFormValue("sex")
//			address = r.PostFormValue("address")
//			cell    = r.PostFormValue("cell")
//		)
//		users, err := services.QueryUser(id, name, sex, address, cell)
//		if err != nil {
//			ErrorMsg(w, err.Error())
//			return
//		}
//		// display
//		t, err := template.New("display").ParseFiles("template/display.html")
//		if err != nil {
//			panic(err)
//		}
//		t.ExecuteTemplate(w, "display.html", users)
//	}
//}
//
func (c *UserController) ErrorMsg(msg string) {
	c.Data["msg"] = msg
	c.TplName = "user/error.html"
}

func (c *UserController) Confirm(msg string) {
	c.Data["msg"] = msg
	c.TplName = "user/confirm.html"
}
