package controllers

import (
	"errors"
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

// Prepare handle user login
func (c *UserController) Prepare() {
	c.Data["cUser"] = nil
	user := c.GetSession("user")
	if user == nil {
		// not logged
		action := c.GetString("action")
		fmt.Println("action:", action)
		if action == "register" {
			c.Redirect(c.URLFor("UserController.Register"), 301)
			return
		}
		//c.Redirect("/auth/login/", 302)
		c.Redirect(c.URLFor("AuthController.Login"), 302)
		return
	}
	if id, ok := user.(int64); ok {
		if user, _ := services.IDFindUser(id); user != nil {
			c.Data["cUser"] = user
		}
	}
	if c.Data["cUser"] == nil {
		c.DestroySession()
		c.Redirect(c.URLFor("AuthController.Login"), 302)
		return
	}
}

// Home give a default page, with a list of users
func (c *UserController) Home() {
	type dUser struct {
		ID       int64
		Name     string
		Sex      int
		Address  string
		Cell     string
		Born     string
		Password string
	}
	cUsers := []dUser{}
	users, err := services.ListAllUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		u := dUser{
			ID:       user.ID,
			Name:     user.Name,
			Sex:      user.Sex,
			Address:  user.Address,
			Cell:     user.Cell,
			Born:     user.Born.Format("2006-01-02"),
			Password: user.Password,
		}
		cUsers = append(cUsers, u)
	}
	c.Data["users"] = cUsers
	c.Data["msg"] = "Are you sure to delete the user?"
	c.TplName = "user/home.html"
}

// Create add a user to user list if request method is post
func (c *UserController) Create() {
	if c.Ctx.Input.IsPost() {
		bornStr := c.GetString("born")
		t, errP := time.Parse("2006.01.02", bornStr)
		if errP != nil {
			HandleError(c, errP)
			return
		}
		sex, err := c.GetInt("sex")
		if err != nil {
			HandleError(c, err)
			return
		}
		name := c.GetString("name")
		address := c.GetString("address")
		cell := c.GetString("cell")
		password := c.GetString("passwd")
		// create user here
		errc := services.CreateUser(name, password, address, cell, sex, t)
		if errc != nil {
			HandleError(c, errc)
			return
		}
		// redirect to home after create user
		c.Redirect(c.URLFor("UserController.Home"), 301)
	} else {
		c.TplName = "user/create.html"
	}
}

// Delete delete a user based on id
func (c *UserController) Delete() {
	id, err := c.GetInt64("id")
	if err != nil {
		HandleError(c, err)
		return
	}
	if id == models.AdminID {
		HandleError(c, errors.New("You can't delete admin, who's id is: "+c.GetString("id")))
		return
	}
	fmt.Println("To delete: ", c.GetString("id"))
	if errd := services.IDDelUser(id); errd == nil {
		HandleError(c, errors.New("Deleted a user, who's id is: "+c.GetString("id")))
		return
	} else {
		HandleError(c, errd)
		return
	}
}

// Edit edit a user by id
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
		if err != nil {
			HandleError(c, err)
		}
		user, errf := services.IDFindUser(id)
		if errf != nil {
			HandleError(c, errf)
		}
		cuser := func(u models.User) cUser {
			return cUser{
				ID:      user.ID,
				Name:    user.Name,
				Sex:     user.Sex,
				Address: user.Address,
				Cell:    user.Cell,
				Born:    strings.Split(user.Born.String(), " ")[0],
				Passwd:  user.Password,
			}
		}(*user)
		c.Data["user"] = cuser
		c.TplName = "user/edit.html"
	} else {
		id, erri := c.GetInt64("id")
		if erri != nil {
			HandleError(c, erri)
		}
		name := c.GetString("name")
		sex := c.GetString("sex")
		address := c.GetString("address")
		cell := c.GetString("cell")
		born := c.GetString("born")
		password := c.GetString("passwd")
		if id == models.AdminID {
			HandleError(c, errors.New("Do not edit admin."))
			return
		}
		if errm := services.IDModUser(name, address, password, cell, sex, born, id); errm != nil {
			HandleError(c, errm)
			return
		} else {
			c.Redirect(c.URLFor("UserController.Home"), 301)
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
		if err != nil {
			HandleError(c, err)
		}
		c.Data["users"] = users
		c.TplName = "user/display.html"
	}
}

func (c *UserController) Detail() {
	user := c.GetString("name")
	fmt.Println("user detail: ", user)
	u, err := services.NameFindUser(user)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.Data["users"] = []*models.User{u}
	c.TplName = "user/display.html"
}

// Register register a user
func (c *UserController) Register() {
	if c.Ctx.Input.IsPost() {
		fmt.Println("creat user...")
	} else {
		c.TplName = "user/create.html"
	}
}

// ResetPass reset a user's pass
func (c *UserController) ResetPass() {
	if c.Ctx.Input.IsPost() {
		name := c.GetString("name")
		oldPass := c.GetString("oldPassword")
		newPass := c.GetString("newPassword")
		confirmPass := c.GetString("confirmPassword")
		if newPass != confirmPass {
			HandleError(c, errors.New("Your newPassword doesn't match confirmPassword"))
			return
		}
		fmt.Printf("name: %#v, oldpass: %#v, newpass: %#v\n", name, oldPass, newPass)
		u, err := services.NameFindUser(name)
		if err != nil {
			HandleError(c, err)
			return
		}
		fmt.Println("user to reset pass: ", u)
		erru := services.UpdateUserPass(u, oldPass, newPass)
		if erru != nil {
			HandleError(c, erru)
			return
		} else {
			c.Redirect(c.URLFor("UserController.Home"), 302)
		}
	} else {
		c.TplName = "user/resetpass.html"
	}
}
