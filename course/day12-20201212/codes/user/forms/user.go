package forms

type AddUserForm struct {
	Name string `form:"name"`
	Addr string `form:"addr"`
	Sex  bool   `form:"sex"`
}
