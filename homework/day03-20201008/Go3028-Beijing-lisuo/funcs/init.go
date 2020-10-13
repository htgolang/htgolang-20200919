package funcs

import (
	"fmt"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/define"
	utils "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/utils"
)

func NewUser(n, p, a string, id int64) map[int64]define.User {
	return map[int64]define.User{
		id: {
			Name:    n,
			Address: a,
			Phone:   p,
		},
	}
}

func AddUser(usersList *[]map[int64]define.User, user map[int64]define.User) {
	define.UserList = append(*usersList, user)
}

func Init() {
	res0 := utils.GenId()
	res1 := utils.GenId()
	fmt.Println(define.Id)
	fmt.Println(res0)
	fmt.Println(res1)
	user0 := NewUser("lisuo", "Beijing", "999", res0)
	AddUser(&define.UserList, user0)
	fmt.Println(define.UserList)
	fmt.Println(define.UserList[0])
	fmt.Println(define.UserList[0][res0])
	fmt.Println(define.UserList[0][res0].Name)
	fmt.Println(define.UserList[0][res0].Address)
	fmt.Println(define.UserList[0][res0].Phone)
}
