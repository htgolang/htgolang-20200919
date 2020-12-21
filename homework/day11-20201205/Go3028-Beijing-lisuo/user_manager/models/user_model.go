package models

import (
	"crypto/md5"
	"fmt"
	"time"
)

// UserQuit represent quit exit status code
const (
	AdminName string = "admin"
	AdminID   int64  = 0
	UserQuit  int    = 1
)

// User to contain user's info
// UserList to contain all the users
type User struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Cell    string    `json:"cell"`
	Born    time.Time `json:"born"`
	Passwd  string    `json:"passwd"`
}

// UserList contains users
var UserList []User

// UserField slice for GetField func
var UserField []string = []string{"Id", "Name", "Address", "Cell", "Born", "Passwd"}

// NewUser make a new user contains user's info
func NewUser(id int64, name, cell, address, born, passwd string) User {
	return User{
		ID:      id,
		Name:    name,
		Cell:    cell,
		Address: address,
		Born: func() time.Time {
			t, _ := time.Parse("2006.01.02", born)
			return t
		}(),
		Passwd: fmt.Sprintf("%x", md5.Sum([]byte(passwd))),
	}
}
