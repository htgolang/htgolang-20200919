package define

import (
	"crypto/md5"
	_ "errors"
	"fmt"
	"time"
)

// User to contain user's info
// UserList to contain all the users
// Id is a int64 number for for each user

type User struct {
	id      int
	name    string
	address string
	cell    string
	born    time.Time
	passwd  [16]uint8
}

// UserList contains users
var UserList []User

// type Operator interface {
// 	AddUser()
// 	DelUser()
// 	ModUser()
// 	RefUser()
// }

func (u User) NameIsEmpty() bool {
	return u.name == ""
}
func (u User) AddressIsEmpty() bool {
	return u.name == ""
}
func (u User) CellIsEmpty() bool {
	return u.cell == ""
}

// make a new user contains user's info
func NewUser(id int, name, cell, address string, born time.Time, passwd [16]uint8) User {
	return User{
		id:      id,
		name:    name,
		cell:    cell,
		address: address,
		born:    born,
		passwd:  passwd,
	}
}

func addUser(ul *[]User, u User) {
	UserList = append(*ul, u)
	fmt.Printf("user %v added\n", u.name)
	//fmt.Printf("currentUserList %v\n", UserList)
}

func Init() {
	user0 := NewUser(0, "lisuo", "18811992299", "HaidianDistrict,BeijingXinParkRestaurants,BeixiaguanSubdistrict,HaidianDistrict,China", time.Now(), md5.Sum([]byte("hello")))
	user1 := NewUser(1, "jack ma", "18834592299", "Venus", time.Now(), md5.Sum([]byte("hi")))
	user2 := NewUser(2, "stevenux", "18821312299", "Jupter", time.Now(), md5.Sum([]byte("hola")))
	user3 := NewUser(3, "jaccyli", "12222392299", "Mars", time.Now(), md5.Sum([]byte("你好")))
	addUser(&UserList, user0)
	addUser(&UserList, user1)
	addUser(&UserList, user2)
	addUser(&UserList, user3)
}
