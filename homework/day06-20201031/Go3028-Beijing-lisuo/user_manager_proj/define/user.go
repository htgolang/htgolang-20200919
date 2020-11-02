package define

import (
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
	ID      int64
	Name    string
	Address string
	Cell    string
	Born    time.Time
	Passwd  string
}

// UserList contains users
var UserList []User

// UserField slice for GetField func
var UserField []string = []string{"Id", "Name", "Address", "Cell", "Born", "Passwd"}
