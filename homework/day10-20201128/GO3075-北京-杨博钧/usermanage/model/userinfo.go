package model

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Sex string `json:"sex"`
	Addr string `json:"addr"`
	Tel string `json:"tel"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}