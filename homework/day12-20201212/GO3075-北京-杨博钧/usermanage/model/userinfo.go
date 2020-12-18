package model

import "time"

type User struct {
	Id int
	Name string
	Sex bool
	Addr string
	Tel string
	Birthday time.Time
	Password string
}

func NewUser(id int, name string, sex bool, addr string, tel string, birthday time.Time, password string) *User {
	return &User{id, name, sex, addr, tel, birthday, password}
}