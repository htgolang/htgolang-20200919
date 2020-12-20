package models

type AddUserForm struct {
	Name string `form:"name"`
	Sex  bool   `form:"sex"`
	Addr string `form:"addr"`
}

type ModifyUserForm struct {
	ID   int64  `form:"id"`
	Name string `form:"name"`
	Sex  bool   `form:"sex"`
	Addr string `form:"addr"`
}
