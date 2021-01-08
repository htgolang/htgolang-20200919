/*
	用户认证处理模块，仅支持账号密码认证
*/

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"user/models"
	"user/utils"
)

type AuthController struct {
	beego.Controller
}

// 登录表单
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Remember string `form:"remember"`
}

// 验证SessionID
func (c *AuthController) VerifySID() bool {
	if v := c.GetSession("user"); v != nil {
		return true
	}
	return false
}

// 验证用户名密码
func (c *AuthController) VerifyUser(name, password string) bool {
	// 验证用户
	ormer := orm.NewOrm()
	user := &models.User{Name: name}
	if err := ormer.Read(user, "Name"); err != nil {
		return false
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}

	return true
}

// 登录失败处理函数
func (c *AuthController) loginError(message string) {
	c.TplName = "login.html"
	c.Data["error"] = message
	return
}

// 登录
func (c *AuthController) Login() {
	// Get请求
	if c.Ctx.Input.IsGet() {
		// 判断sessionID
		if c.VerifySID() {
			c.Redirect("/", 302)
		} else {
			c.TplName = "login.html"
		}
	}

	// Post
	if c.Ctx.Input.IsPost() {
		// 解析表单
		form := &LoginForm{}
		if err := c.ParseForm(form); err != nil {
			c.loginError("服务器内部错误")
			return
		}

		// 验证用户名密码
		if c.VerifyUser(form.Username, form.Password) {
			user, _ := utils.GetUserByName(form.Username)
			c.SetSession("user", user.ID)

			/*
				下次自动登录功能
				如果勾选，将Cookie过期时间设置和Session一致，
				如果未选，将Cookie过期时间设置为0
			*/
			cookieName := beego.BConfig.WebConfig.Session.SessionName
			cookieValue := c.CruSession.SessionID()
			cookieMaxAge := beego.BConfig.WebConfig.Session.SessionGCMaxLifetime

			if form.Remember == "on" {
				c.Ctx.SetCookie(cookieName, cookieValue, cookieMaxAge)
			} else {
				c.Ctx.SetCookie(cookieName, cookieValue, 0)
			}
		} else {
			c.loginError("用户名或密码错误")
			return
		}

		// 重定向到后台首页
		c.Redirect("/", 302)
	}
}

func (c *AuthController) Logout() {
	// 删除Session内全部键值对
	//c.DestroySession()
	// 删除Session内指定键值对
	c.DelSession("user")
	c.Redirect("/auth/login", 302)
}
