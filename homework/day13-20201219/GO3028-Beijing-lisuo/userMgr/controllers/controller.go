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

func (c *UserController) Prepare() {
	user := c.GetSession("user")
	if user == nil {
		// not logged
		c.Redirect("/auth/login/", 302)
	}
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
		r := c.Ctx.Request.Referer()
		bornStr := c.GetString("born")
		t, errP := time.Parse("2006.01.02", bornStr)
		if errP != nil {
			c.ErrorMsg(errP.Error(), r)
			return
		}
		sex, err := c.GetInt("sex")
		if err != nil {
			c.ErrorMsg(err.Error(), r)
			return
		}
		name := c.GetString("name")
		address := c.GetString("address")
		cell := c.GetString("cell")
		password := c.GetString("passwd")

		// create user here
		if errc := services.CreateUser(name, password, address, cell, sex, t); errc != nil {
			c.ErrorMsg(errc.Error(), r)
			return
		}

		// redirect to home after create user
		c.Redirect("/user/home/", 301)
	} else {
		c.TplName = "user/create.html"
	}
}

// Delete delete a user based on id
func (c *UserController) Delete() {
	r := c.Ctx.Request.Referer()
	id, err := c.GetInt64("id")
	if err != nil {
		c.ErrorMsg(err.Error(), r)
		return
	}
	if id == models.AdminID {
		c.ErrorMsg("You can't delete admin, who's id is: "+c.GetString("id"), r)
		return
	}
	fmt.Println("To delete: ", c.GetString("id"))
	if errd := services.IDDelUser(id); errd == nil {
		c.ErrorMsg("Deleted a user, who's id is: "+c.GetString("id"), r)
		return
	} else {
		HandleError(c, errd, r)
	}
}

// Edit edit a user by id
func (c *UserController) Edit() {
	if c.Ctx.Input.IsGet() {
		r := c.Ctx.Request.Referer()
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
		if err != nil {
			HandleError(c, err, r)
		}
		user, errf := services.IDFindUser(id)
		if errf != nil {
			HandleError(c, errf, r)
		}
		cuser := func(models.User) cUser {
			return cUser{
				ID:      user.ID,
				Name:    user.Name,
				Sex:     user.Sex,
				Address: user.Address,
				Cell:    user.Cell,
				Born:    strings.Split(user.Born.String(), " ")[0],
				Passwd:  user.Password,
			}
		}(user)
		c.Data["user"] = cuser
		c.TplName = "user/edit.html"
	} else {
		r := c.Ctx.Request.Referer()
		id, erri := c.GetInt64("id")
		if erri != nil {
			HandleError(c, erri, r)
		}
		name := c.GetString("name")
		sex := c.GetString("sex")
		address := c.GetString("address")
		cell := c.GetString("cell")
		born := c.GetString("born")
		password := c.GetString("passwd")
		if id == models.AdminID {
			c.ErrorMsg("Do not edit admin.", r)
			return
		}
		if errm := services.IDModUser(name, address, password, cell, sex, born, id); errm != nil {
			HandleError(c, errm, r)
			return
		} else {
			c.Redirect("/user/home/", 301)
		}
	}
}

// Query get users by name, address, cell or id
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
		fmt.Println("users and err in services.QueryUser", users, err)
		r := c.Ctx.Request.Referer()
		if err != nil {
			HandleError(c, err, r)
		}
		c.Data["users"] = users
		c.TplName = "user/display.html"
	}
}

// HandleError wrap err handle code
func HandleError(c *UserController, err error, refer string) {
	if err != nil {
		c.ErrorMsg(err.Error(), refer)
		return
	}
}

// ErrorMsg send errors to client
func (c *UserController) ErrorMsg(msg, r string) {
	c.Data["msg"] = msg
	c.Data["refer"] = r
	c.TplName = "user/error.html"
}
