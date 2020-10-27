package models

import (
	"time"
)

//User ...
//Global data structure of User element.
type User struct {
	ID       int
	Name     string
	Tel      string
	Address  string
	Birthday time.Time
	Password [16]byte
}

//Users ...
//Global Data Structure, for each element of Users slice,
//in other words, User, has the same data structure below:
/*
	ID int
	Name string
	Tel string
	Address string
	Birthday time.Time
	password [16]byte
*/
var Users []User

//LoginCount ...
// Login count number
var LoginCount int = 0

func init() {
	Users = make([]User, 0)
	Users = append(Users, *GenerateElement(1001, "admin", "+1 4406665321", "2426 Wildwood Street, Medina, Ohio", "1990-01-01", "admin"))
	Users = append(Users, *GenerateElement(1002, "root", "+1 4406665322", "2427 Wildwood Street, Medina, Ohio", "1992-01-01", "root"))
	Users = append(Users, *GenerateElement(1003, "super", "+1 4406665323", "2428 Wildwood Street, Medina, Ohio", "1994-01-01", "super"))
}
