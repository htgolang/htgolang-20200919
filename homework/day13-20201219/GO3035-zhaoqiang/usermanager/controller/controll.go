package controller

import (
	"fmt"
	"user/models"
	"user/server"
)

type UserController struct {
	BaseController
}

func (u *UserController) Jump() {
	u.Redirect("/user/list", 301)
}

func (u *UserController) List() {
	users, err := server.GetUsers()
	u.Data["error"] = err
	u.Data["users"] = users
	u.TplName = "list.html"
}

func (u *UserController) Create() {
	var errmsg error
	// if u.username != true {
	// 	u.Ctx.Output.Body([]byte("没有权限"))
	// 	return
	// }

	if u.Ctx.Input.IsPost() {
		var user models.User //注意创建指针类型需要使用new函数去创建
		errmsg = u.ParseForm(&user)
		errmsg = server.AuthName(&user, server.CreateUser)
		// fmt.Println("----->", errmsg)
		if errmsg == nil {
			u.Redirect("/user/list", 301)
		}
	}
	u.Data["error"] = errmsg
	u.TplName = "create.html"
}
func (u *UserController) Delete() {
	var id int64
	err := u.Ctx.Input.Bind(&id, "id")
	if err != nil {
		u.Ctx.Output.SetStatus(500)
		u.Ctx.Output.Body([]byte("删除用户,解析id错误"))
	}
	if id == 0 || id == 1 {
		u.Redirect("/user/list", 301)
	}
	err = server.DeleteUserByID(id)
	if err != nil {
		u.Ctx.Output.SetStatus(500)
		u.Ctx.Output.Body([]byte(fmt.Sprintf("删除用户,删除错误,id=%d", id)))
	}
	u.Redirect("/user/list", 301)
}

func (u *UserController) Modify() {
	var errmsg error
	var id int64
	errmsg = u.Ctx.Input.Bind(&id, "id")
	if id == 0 {
		u.Ctx.Output.SetStatus(500)
		u.Ctx.Output.Body([]byte("修改用户,解析id错误"))
	}
	user, err := server.QueryUserByID(id)
	if err != nil {
		u.Ctx.Output.SetStatus(500)
		u.Ctx.Output.Body([]byte(fmt.Sprintf("查询用户错误,id=%d", id)))
	}
	brith := user.Brithday.Format("2006-01-02")

	if u.Ctx.Input.IsPost() {
		newuser := new(models.User)
		err := u.ParseForm(newuser)
		newuser.CreateAt = user.CreateAt
		// fmt.Println(newuser)///////
		if err != nil {
			errmsg = err
		} else {
			err := server.AuthName(newuser, server.ModifyUser)
			if err != nil {
				errmsg = err
			} else {
				u.Redirect("/user/list", 301)
			}
		}

	}
	u.Data["error"] = errmsg
	u.Data["user"] = user
	u.Data["brith"] = brith
	u.TplName = "modify.html"
}
