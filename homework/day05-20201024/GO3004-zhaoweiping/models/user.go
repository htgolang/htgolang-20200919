package models

import (
	"fmt"
	"strings"
	"time"
)

// var UsersDb = []Users{}

var UsersDb = make(map[string]*Users)

type Users struct {
	id       int
	name     string
	addr     string
	tel      string
	birthday time.Time
	passwd   string
}

// var UserDb []Users

func NewUser(id int, name string, addr string, tel string, birthday time.Time, passwd string) *Users {
	// return &Users{id, name, addr, tel, birthday, passwd}
	return &Users{
		id:       id,
		name:     name,
		addr:     addr,
		tel:      tel,
		birthday: birthday,
		passwd:   passwd,
	}
}

func GetId() int {
	id := 0
	for _, user := range UsersDb {
		if user.id > id {
			id = user.id
		}
	}
	return id + 1
}

func (u *Users) AddUser() {
	u.id = GetId()
	for _, user := range UsersDb {
		if user.name == u.name {
			fmt.Println("名称（name）已存在，请重新输入莫重复！！！")
			break
		}
	}
	UsersDb[u.name] = u

	fmt.Printf("用户【%v】添+成功！！！\n", u.name)
}

func (u *Users) FindUserById(id int) *Users {
	for _, user := range UsersDb {
		if user.id == id {
			return user
		}
	}
	return nil
}

func (u *Users) ModifyUserById() {
	for _, tuser := range UsersDb {
		if tuser.id == u.id {
			if tuser.name == u.name {
				fmt.Println("输入的名称已经存在，请重新输入")
				break
			} else {
				delete(UsersDb, tuser.name)
				UsersDb[u.name] = u
				fmt.Println("用户修改完成。")
				u.PrintUser()
			}
		}
	}
}

func (u *Users) DeleteUserById() {
	delete(UsersDb, u.name)
}

func filter(user *Users, q string) bool {
	return strings.Contains(user.name, q) ||
		strings.Contains(user.addr, q) ||
		strings.Contains(user.tel, q)
}

func QueryUser(q string) map[string]*Users {
	rt := make(map[string]*Users)
	for _, user := range UsersDb {
		// fmt.Println(user)
		if filter(user, q) {
			rt[user.name] = user
		}
	}
	return rt
}

func PrintUsersDb() {
	fmt.Println("\n以下是目前已存在的用户：")
	for _, user := range UsersDb {
		fmt.Printf("ID：%v, 名称：%v, 联系方式：%v, 通信地址：%v，生日：%v\n", user.id, user.name, user.tel, user.addr, user.birthday)
	}
}

func (u *Users) PrintUser() {
	fmt.Printf("ID：%v, 名称：%v, 联系方式：%v, 通信地址：%v，生日：%v\n", u.id, u.name, u.tel, u.addr, u.birthday)
}
