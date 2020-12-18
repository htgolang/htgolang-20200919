package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"usermanage/model"
	"usermanage/services"
	"usermanage/utils"
	"usermanage/forms"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) AddPage() {
	form := model.NewAddUpdatePage()
	if c.Ctx.Input.IsGet() {
		c.Data["form"] = form
		c.TplName = "user/addpage.html"
	} else {
		var Data forms.User
		c.ParseForm(&Data)
		err := services.ParseAddUpdateParams(Data, form)
		if utils.IsAddUserExists(Data.Name){
			form.NameError = "当前用户名已存在,不能进行添加!"
			form.PasswdError = ""
			form.BirthdayError = ""
			c.Data["form"] = form
			c.TplName = "user/addpage.html"
		} else if !utils.IsPasswdSame(Data){
			form.NameError = ""
			form.PasswdError = "两次输入的密码不一致请重新输入"
			form.BirthdayError = ""
			c.Data["form"] = form
			c.TplName = "user/addpage.html"
		} else if err != nil {
			form.NameError = ""
			form.PasswdError = ""
			form.BirthdayError = fmt.Sprintf("%v", err)
			c.Data["form"] = form
			c.TplName = "user/addpage.html"
		} else {
			form.NameError = ""
			form.PasswdError = ""
			form.BirthdayError = ""
			services.AddUser(form)
			c.Redirect("/user/mainpage", 302)
		}
	}
}

func (c *UserController) MainPage() {

	mainpage := model.NewMainPage()
	if c.Ctx.Input.IsGet() {
		mainpage.Userinfos = services.GetAllUser()
		c.Data["form"] = mainpage
		c.TplName = "user/mainpage.html"
	} else {
		var Data forms.QueryInfo
		c.ParseForm(&Data)
		err := services.ParseQueryParams(Data, mainpage)
		if err != nil {
			mainpage.Error = fmt.Sprintf("%v", err)
			mainpage.Userinfos = services.GetAllUser()
			c.Data["form"] = mainpage
			c.TplName = "user/mainpage.html"
		} else {
			mainpage.Error = ""
			mainpage.Userinfos = services.GetQueryUser(mainpage)
			c.Data["form"] = mainpage
			c.TplName = "user/mainpage.html"
		}
	}
}

func (c *UserController) UpdatePage() {
	updatepage := model.NewAddUpdatePage()
	if c.Ctx.Input.IsGet() {
		updatepage.NameError = ""
		updatepage.PasswdError = ""
		updatepage.BirthdayError = ""
		id, _ := c.GetInt("Id")
		services.GetPageById(id, updatepage)
		updatepage.Id = id
		c.Data["form"] = updatepage
		c.TplName = "user/updatepage.html"
	} else {
		var Data forms.User
		c.ParseForm(&Data)
		err := services.ParseAddUpdateParams(Data, updatepage)
		id, _ := c.GetInt("Id")
		if utils.IsUpdateUserExists(id, Data.Name){
			updatepage.NameError = "当前用户名已存在,不能进行添加!"
			updatepage.PasswdError = ""
			updatepage.BirthdayError = ""
			c.Data["form"] = updatepage
			c.TplName = "user/updatepage.html"
		} else if !utils.IsPasswdSame(Data){
			updatepage.NameError = ""
			updatepage.PasswdError = "两次输入的密码不一致请重新输入"
			updatepage.BirthdayError = ""
			c.Data["form"] = updatepage
			c.TplName = "user/updatepage.html"
		} else if err != nil {
			updatepage.NameError = ""
			updatepage.PasswdError = ""
			updatepage.BirthdayError = fmt.Sprintf("%v", err)
			c.Data["form"] = updatepage
			c.TplName = "user/updatepage.html"
		} else {
			updatepage.NameError = ""
			updatepage.PasswdError = ""
			updatepage.BirthdayError = ""
			services.UpdateUser(id, updatepage)
			c.Redirect("/user/mainpage", 302)
		}
	}
}

func (c *UserController) DeletePage() {
	id, _ := c.GetInt("Id")
	services.DeleteUser(id)
	c.Redirect("/user/mainpage", 302)
}