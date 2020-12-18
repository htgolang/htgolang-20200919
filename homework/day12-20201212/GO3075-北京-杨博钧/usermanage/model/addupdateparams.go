package model

type AddUpdatePage struct {
	NameError string
	PasswdError string
	BirthdayError string
	Id int
	Name string
	IsMale string
	IsFeMale string
	Addr string
	Tel string
	Birthday string
	Passwd string
}

func NewAddUpdatePage() *AddUpdatePage {
	return &AddUpdatePage{IsMale: "checked",}
}
