package controllers

import (
	"CMS/models"
	"CMS/services"
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Index() {
	// c.Ctx.WriteString("nihao")
	c.TplName = "user/add.html"
	// c.Redirect("/user/query", 302)
}

func (c *UserController) Add() {
	if c.Ctx.Input.IsPost() {
		var form models.AddUserForm
		c.ParseForm(&form)
		services.AddUser(form.Name, form.Sex, form.Addr)
		c.Redirect("/user/query", 302)
	} else {
		c.Data["form"] = &models.AddUserForm{Name: "kk"}
		c.TplName = "user/add.html"

	}
}

func (c *UserController) Query() {
	c.Data["users"] = services.GetUsers()
	c.TplName = "user/query.html"
}

func (c *UserController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		services.DeleteUser(id)
	}
	c.Redirect("/user/query", 302)
}

var a string

func (c *UserController) Modify() {
	// id := c.Input().GetInt64("id")
	id, err := c.GetInt64("id")
	if err != nil {
		fmt.Println(err)
	}
	if c.Ctx.Input.IsPost() {
		var form models.ModifyUserForm
		c.ParseForm(&form)
		services.ModifyUser(id, form.Name, form.Sex, form.Addr)
		c.Redirect("/user/query", 302)
	} else {
		c.Data["form"] = &models.ModifyUserForm{ID: id}
		c.TplName = "user/modify.html"

	}
}
