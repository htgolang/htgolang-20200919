package controllers

import (
	"cmdb/errors"
	"cmdb/forms"
	"cmdb/services"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Login() {
	if user := c.GetSession("user"); user != nil {
		c.Redirect(beego.URLFor("UserController.Query"), 302)
		return
	}
	// 打开
	// 点击登陆
	form := &forms.LoginForm{}
	// 错误信息 attr = []error
	errors := errors.NewErrors()

	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {

			valid := validation.Validation{}

			valid.Required(form.Username, "username.username.username").Message("用户名不能为空")
			valid.Required(form.Password, "password.password.password").Message("密码不能为空")

			if !valid.HasErrors() {
				if user := services.Auth(form.Username, form.Password); user != nil {
					fmt.Println("登陆成功")
					c.SetSession("user", user.ID)
					c.Redirect(beego.URLFor("UserController.Query"), 302)
					return
				} else {
					errors.AddError("default", "用户名或密码错误")
				}
			} else {
				// k => []Error
				// k => []String
				errors.AddErrorMap(valid.ErrorMap())
			}
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	fmt.Println(beego.BConfig.WebConfig.XSRFKey)
	fmt.Println(beego.BConfig.WebConfig.XSRFExpire)
	fmt.Println(c.XSRFExpire)
	fmt.Println(c.XSRFToken())
	c.Data["errors"] = errors
	c.TplName = "auth/login.html"
}

func (c *AuthController) Logout() {
	c.DestroySession()
	c.Redirect(beego.URLFor("AuthController.Login"), 302)
}
