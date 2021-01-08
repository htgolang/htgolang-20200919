package models

type LoginUser struct {
	ID          int64 `form:"id"`
	Name        string `form:"name"`
	PassWord    string `form:"password"`
	NewPassWord string `form:"newpassword"`
}
