package models

type LoginUser struct {
	Name     string `form:"name"`
	PassWord string `form:"password"`
}


