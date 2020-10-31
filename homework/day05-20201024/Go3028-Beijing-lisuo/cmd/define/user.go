package define

import (
	"time"
)

const USER_QUIT = 1

// User to contain user's info
// UserList to contain all the users
// Id is a int64 number for for each user
// Passwd is a [16]uint8 type for saving
type User struct {
	ID      int64
	Name    string
	Address string
	Cell    string
	Born    time.Time
	Passwd  [16]byte
}

// UserList contains users
var UserList []User

// type Operator interface {
// 	AddUser()
// 	DelUser()
// 	ModUser()
// 	RefUser()
// }

// UserField slice for GetField func
var UserField []string = []string{"Id", "Name", "Address", "Cell", "Born", "Passwd"}

//
//func (u User) NameIsEmpty() bool {
//	return u.Name == ""
//}
//func (u User) AddressIsEmpty() bool {
//	return u.Name == ""
//}
//func (u User) CellIsEmpty() bool {
//	return u.Cell == ""
//}
