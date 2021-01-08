package controller

import (
	"errors"
	"user/server"
)

type PermissionController struct {
	BaseController
}

func (p *PermissionController) Prepare() {
	p.BaseController.Prepare()
	p.Data["navkey"] = ""
}

func (p *PermissionController) Setting() {
	var errmsg error
	var id int64

	err := p.Ctx.Input.Bind(&id, "id")
	if err != nil {
		errmsg = err
	}

	user, err := server.QueryUserByID(id)
	if err != nil {
		errmsg = err
	}
	if !p.userinfo.Admin {
		errmsg = errors.New("用户没有权限")
	} else {
		if p.Ctx.Input.IsPost() {
			var permiss bool = false
			err := p.Ctx.Input.Bind(&permiss, "permiss")
			if err != nil {
				errmsg = err
			} else {
				user.Admin = permiss
				err := server.ModifyUserPerm(user)
				if err != nil {
					errmsg = err
				} else {
					p.Redirect("/user/list", 301)
				}

			}
		}
	}
	p.Data["error"] = errmsg
	p.Data["user"] = user
	p.TplName = "permission.html"
}
