package controllers

import (
	base "cmdb/base/controllers"
	"cmdb/forms"
	"cmdb/services"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	base.LayoutController
}

func (c *UserController) Add() {
	var form forms.AddUserForm
	valid := validation.Validation{}

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(&form); err == nil {
			if success, err := valid.Valid(&form); err == nil && success {
				services.AddUser(form.Name, form.Password, form.Addr, form.Sex)
				flash := beego.NewFlash()
				flash.Set("success", "新建成功")
				flash.Store(&c.Controller)
				c.Redirect(beego.URLFor("UserController.Query"), 302)
				return
			}
		}
	}
	c.Data["errors"] = valid.ErrorMap()
	c.Data["form"] = &form

	c.TplName = "user/add.html" // LayoutContent内容替代的文件路径
	c.Layout = "base/layout.html"
	c.LayoutSections["LayoutSectionTitle"] = "user/add_title.html"
}

func (c *UserController) Query() {
	c.Data["users"] = services.GetUsers()

	c.TplName = "user/query.html"
	c.LayoutSections["LayoutSectionTitle"] = "user/query_title.html"
}

func (c *UserController) QueryJson() {
	c.Data["json"] = services.GetUsers()
	c.ServeJSON()
}

func (c *UserController) QueryXml() {
	c.Data["xml"] = services.GetUsers()
	c.ServeXML()
}

func (c *UserController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		services.DeleteUser(id)
		flash := beego.NewFlash()
		flash.Set("success", "删除成功")
		flash.Store(&c.Controller)
	}
	c.Redirect(beego.URLFor("UserController.Query"), 302)
}
