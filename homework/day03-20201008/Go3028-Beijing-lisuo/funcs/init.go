package funcs

import (
	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/define"
)

func NewUser(n, p, a string) define.User {
	return define.User{
		Name:    n,
		Phone:   p,
		Address: a,
	}
}

func Init() {
	user0 := NewUser("lisuo", "18811738208", "HaidianDistrict,BeijingXinParkRestaurants,BeixiaguanSubdistrict,HaidianDistrict,China")
	user1 := NewUser("jaccyli", "13899998888", "London,England,Earth")
	user2 := NewUser("stevenux", "12899998889", "London,England,Earth")
	AddUser(&define.UserList, user0)
	AddUser(&define.UserList, user1)
	AddUser(&define.UserList, user2)
	//fmt.Println(define.UserList)
	ShowUserList()
	//fmt.Println(define.UserList[0])
	//fmt.Println(define.UserList[0][res0])
	//fmt.Println(define.UserList[0][res0].Name)
	//fmt.Println(define.UserList[0][res0].Address)
	//fmt.Println(define.UserList[0][res0].Phone)
}
