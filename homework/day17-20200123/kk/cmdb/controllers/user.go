package controllers

import (
	base "cmdb/base/controllers"
	"cmdb/forms"
	"cmdb/services"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	base.RequiredAuthController
}

func (c *UserController) Prepare() {
	c.RequiredAuthController.Prepare()
	c.Data["navKey"] = "user"
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
	c.LayoutSections = make(map[string]string)
	// map key LayoutSectionXXX
	// 		value 需要填充内容的html文件路径
	c.LayoutSections["LayoutSectionTitle"] = "user/add_title.html"
}

func (c *UserController) Query() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["users"] = services.GetUsers()
	c.TplName = "user/query.html"
	c.Layout = "base/layout.html"
	c.LayoutSections = make(map[string]string)
	// map key LayoutSectionXXX
	// 		value 需要填充内容的html文件路径
	c.LayoutSections["LayoutSectionTitle"] = "user/query_title.html"
}

func (c *UserController) QueryJson() {

	c.Data["json"] = services.GetUsers()
	c.ServeJSON()
	// yaml
	// Data["yaml"]
	// ServeYAML
	// xml
	// Data["xml"]
	// ServeXML
	// jsonp
	// callback(json)
	// Data["jsonp"]
	//ServeJSONP
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
		// flash.Success("删除成功")
		// flash.Error("删除失败") key => error
		// flash.Warning() key = warning
		flash.Store(&c.Controller)
	}
	c.Redirect(beego.URLFor("UserController.Query"), 302)
}
