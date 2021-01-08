package controller

import (
	"errors"
	"user/models"
	"user/server"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (l *AuthController) Login() {
	if s := l.GetSession("id"); s != nil { //获取到session信息跳转登陆
		l.Redirect("/user/list", 301)
		return
	}

	user := new(models.LoginUser)
	var errmsg error
	err := l.ParseForm(user)
	if err != nil {
		errmsg = err
	}
	if l.Ctx.Input.IsPost() {
		if id, err := server.AuthLogin(user); err != nil {
			errmsg = errors.New("用户名密码不正确")
		} else {
			l.SetSession("id", id)
			l.Redirect("/user/list", 301)
		}
	}

	l.Data["error"] = errmsg
	l.TplName = "login.html"
}

func (l *AuthController) Logout() {
	l.DestroySession()
	l.Redirect("/auth/login", 301)
}
