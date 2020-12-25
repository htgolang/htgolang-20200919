package controller

import (
	"user/server"
)

type PermissionController struct {
	BaseController
}

func (p *PermissionController) Setting() {
	var errmsg error
	var id int64

	session := p.GetSession("logstatu")
	if session != true {
		p.Redirect("/user/list", 301)
	}

	err := p.Ctx.Input.Bind(&id, "id")
	if err != nil {
		errmsg = err
	}

	user, err := server.QueryUserByID(id)
	if err != nil {
		errmsg = err
	}
	// fmt.Println("=====>", user)
	if p.Ctx.Input.IsPost() {
		var permiss bool = false
		err := p.Ctx.Input.Bind(&permiss, "permiss")
		if err != nil {
			errmsg = err
		} else {
			// fmt.Println("======", permiss)
			user.Admin = permiss
			err := server.ModifyUserPerm(user)
			if err != nil {
				errmsg = err
			} else {
				p.Redirect("/user/list", 301)
			}

		}
	}
	p.Data["error"] = errmsg
	p.Data["user"] = user
	p.TplName = "permission.html"
}
