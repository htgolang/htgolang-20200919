package controller

import (
	"user/models"
	"user/server"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	userinfo *models.User
}

func (base *BaseController) Prepare() {
	session := base.GetSession("id")
	if session == nil {
		base.Redirect("/auth/login", 301)
		return
	}
	base.Data["currentuser"] = nil
	// fmt.Println(reflect.TypeOf(session).Kind())
	if a, ok := session.(int64); ok {
		// fmt.Printf("%T", a)  ////////////////
		if user, err := server.QueryUserByID(a); err == nil {
			base.Data["currentuser"] = user
			base.userinfo = user
		}
	}
	if base.Data["currentuser"] == nil { //断言失败，或者获取用户失败
		base.DestroySession()
		base.Redirect("/auth/login", 301)
	}
	// base.username = *(*string)(unsafe.Pointer(&session))

}
