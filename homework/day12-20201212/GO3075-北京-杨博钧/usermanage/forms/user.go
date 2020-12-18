package forms

type User struct {
	Name string `form:"name"`
	Sex bool `form:"sex"`
	Addr string `form:"addr"`
	Tel string `form:"phone"`
	Birthday string `form:"birthday"`
	Password string `form:"password"`
	Confirm string `form:"confirm"`
}

type QueryInfo struct {
	Id string `form:"id"`
	Name string `form:"name"`
	Sex string `form:"sex"`
	Addr string `form:"addr"`
	Tel string `form:"phone"`
	Birthday string `form:"birthday"`
}