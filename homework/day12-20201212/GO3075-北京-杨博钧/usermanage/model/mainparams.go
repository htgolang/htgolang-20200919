package model

type MainPage struct {
	Id string
	Name string
	Sex string
	Addr string
	Tel string
	Birthday string
	Error string
	Userinfos []*User
}

func NewMainPage() *MainPage {
	return &MainPage{}
}