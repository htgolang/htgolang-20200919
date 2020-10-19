package funcs

import (
	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day04-20201017/Go3028-Beijing-lisuo/define"
)

// make a new user contains user's info
func NewUser(n, p, a string) define.User {
	return define.User{
		Name:    n,
		Phone:   p,
		Address: a,
	}
}

// add three users to define.Userlist
func Init() {
	user0 := NewUser("lisuo", "18811738208", "HaidianDistrict,BeijingXinParkRestaurants,BeixiaguanSubdistrict,HaidianDistrict,China")
	user1 := NewUser("jaccyli", "13899998888", "London,England,Earth")
	user2 := NewUser("stevenux", "12899998889", "London,England,Earth")
	AddUser(&define.UserList, user0)
	AddUser(&define.UserList, user1)
	AddUser(&define.UserList, user2)
	ShowUserList()
}
