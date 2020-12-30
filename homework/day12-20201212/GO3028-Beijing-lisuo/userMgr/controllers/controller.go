package controllers

import (
	"fmt"
	"strings"
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
	c.Data["msg"] = "Are you sure to delete the user?"
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
	if errd := services.IDDelUser(id); errd == nil {
		c.ErrorMsg("Deleted a user, who's id is: " + c.GetString("id"))
		return
	} else {
		HandleError(c, errd)
	}
}

func (c *UserController) Edit() {
	if c.Ctx.Input.IsGet() {
		type cUser struct {
			ID      int64
			Name    string
			Sex     int
			Address string
			Cell    string
			Born    string
			Passwd  string
		}
		id, err := c.GetInt64("id")
		fmt.Println(c.GetInt64("id"))
		HandleError(c, err)
		user, errf := services.IDFindUser(id)
		HandleError(c, errf)
		cuser := func(models.User) cUser {
			return cUser{
				ID:      user.ID,
				Name:    user.Name,
				Sex:     user.Sex,
				Address: user.Address,
				Cell:    user.Cell,
				Born:    strings.Split(user.Born.String(), " ")[0],
				Passwd:  user.Passwd,
			}
		}(user)
		c.Data["user"] = cuser
		c.TplName = "user/edit.html"
	} else {
		id, erri := c.GetInt64("id")
		HandleError(c, erri)
		name := c.GetString("name")
		sex := c.GetString("sex")
		address := c.GetString("address")
		cell := c.GetString("cell")
		born := c.GetString("born")
		password := c.GetString("passwd")
		if id == models.AdminID {
			c.ErrorMsg("Do not edit admin.")
			return
		}
		if errm := services.IDModUser(name, address, password, cell, sex, born, id); errm != nil {
			HandleError(c, errm)
			return
		} else {
			c.Redirect("/user/home/", 301)
		}
	}
}

func (c *UserController) Query() {
	if c.Ctx.Input.IsGet() {
		c.TplName = "user/query.html"
	} else {
		id := c.GetString("id")
		name := c.GetString("name")
		address := c.GetString("address")
		cell := c.GetString("cell")
		fmt.Printf("query input: id: %#v, name: %#v, address: %#v, cell: %#v\n", id, name, address, cell)
		users, err := services.QueryUser(id, name, address, cell)
		HandleError(c, err)
		c.Data["users"] = users
		c.TplName = "user/display.html"
	}
}

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

func HandleError(c *UserController, err error) {
	if err != nil {
		c.ErrorMsg(err.Error())
		return
	}
}
