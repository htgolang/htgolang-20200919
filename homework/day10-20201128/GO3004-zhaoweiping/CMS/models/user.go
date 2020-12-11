package models

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

// var UsersDb = []Users{}
var uTemp string

var UsersDb = make(map[string]*Users)

type Users struct {
	ID   int
	Name string
	Sex  bool
	Addr string
}

var UserDb []*Users

func NewUser(id int, name string, sex bool, addr string) *Users {
	return &Users{id, name, sex, addr}
	// return &Users{
	// 	ID:   id,
	// 	Name: name,
	// 	sex:  sex,
	// 	Addr: addr,
	// }
}

func GetUsers() []*Users {
	users := make([]*Users, 0, 10)
	for _, user := range UsersDb {
		users = append(users, user)
	}
	return users
}

func GetId() int {
	id := 0
	for _, user := range UsersDb {
		if user.ID > id {
			id = user.ID
		}
	}
	return id + 1
}

func AddUser(name string, sex bool, addr string) bool {
	if !QueryName(name) {
		id := GetId()
		user := NewUser(id, name, sex, addr)
		UsersDb[name] = user
		fmt.Printf("用户【%v】添+成功！！！\n", name)
		return true
	} else {
		fmt.Printf("用户【%v】添+失败！！！\n", name)
		return false
	}
}

func DeleteUser(id int) {
	for k, v := range UsersDb {
		if v.ID == id {
			delete(UsersDb, k)
		}
	}
}

func filter(user *Users, q string) bool {
	return strings.Contains(user.Name, q)
}

func QueryName(q string) bool {
	for _, user := range UsersDb {
		if filter(user, q) {
			return true
		}
	}
	return false
}

func ModifyUserGet(id string) *Users {
	var u *Users
	for _, v := range UsersDb {
		if v.ID == cast.ToInt(id) {
			u = v
		}
	}
	// fmt.Println(u)
	uTemp = u.Name
	return u
}

func ModifyUserPost(id string, name string, sex bool, addr string) bool {
	if name != uTemp {
		fmt.Printf("姓名【%v】不可修改！！！\n", name)
		return false
	} else {
		fmt.Println(UsersDb[name].Sex)
		// user := NewUser(cast.ToInt(id), name, sex, addr)
		UsersDb[name].Sex = sex
		UsersDb[name].Addr = addr
		fmt.Printf("用户【%v】修改成功！！！\n", name)
		return true
	}
}
