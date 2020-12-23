package models

import (
	"crypto/md5"
	"fmt"
	"time"
)

// UserQuit represent quit exit status code
const (
	AdminName string = "admin"
	AdminID   int64  = 5
	UserQuit  int    = 1
)

// User to contain user's info
// UserList to contain all the users
type User struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Sex     int       `json:"sex"`
	Address string    `json:"address"`
	Cell    string    `json:"cell"`
	Born    time.Time `json:"born"`
	Passwd  string    `json:"passwd"`
}

// UserList contains users
type UserList []User

// NewUser make a new user contains user's info
func NewUser(id int64, sex int, name, cell, address, passwd string, born time.Time) User {
	return User{
		ID:      id,
		Name:    name,
		Sex:     sex,
		Cell:    cell,
		Address: address,
		Born:    born,
		Passwd:  fmt.Sprintf("%x", md5.Sum([]byte(passwd))),
	}
}
