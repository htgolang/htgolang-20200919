package define

// User to contain user's info
// UserList to contain all the users
// Id is a int64 number for for each user, contains 12 digits

var Id int64

type User struct {
	Name    string
	Address string
	Phone   string
}

// map id to User
var UserList []map[int64]User

/*
[map[450001057600:{lisuo 999 Beijing}] map[450001057600:{lisuo 999 Beijing}] map[450001057600:{lisuo 999 Beijing}]]
*/
