/*
	用户管理模块
*/

package controllers

import (
	"user/utils"
)

type AdminController struct {
	RequireAuth
}

func (c *AdminController) Get() {
	// 获取当前用户名
	if uid, ok := c.GetSession("user").(int64); ok {
		if user, err := utils.GetUserByID(uid); err == nil {
			c.Data["user"] = user
		}
	}

	// 获取所有用户
	if users, err := utils.GetUsers(); err == nil {
		c.Data["users"] = users
	}

	c.TplName = "admin.html"
}
