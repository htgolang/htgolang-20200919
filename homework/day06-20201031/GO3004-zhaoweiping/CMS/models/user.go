package models

import (
	"fmt"
	"strings"
	"time"
)

// var UsersDb = []Users{}

var UsersDb = make(map[string]*Users)

type Users struct {
	ID       int
	Name     string
	Addr     string
	Tel      string
	Birthday time.Time
	Passwd   string
}

var UserDb []*Users

func NewUser(id int, name string, addr string, tel string, birthday time.Time, passwd string) *Users {
	// return &Users{id, name, addr, tel, birthday, passwd}
	return &Users{
		ID:       id,
		Name:     name,
		Addr:     addr,
		Tel:      tel,
		Birthday: birthday,
		Passwd:   passwd,
	}
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

func (u *Users) AddUser() {
	u.ID = GetId()
	for _, user := range UsersDb {
		if user.Name == u.Name {
			fmt.Println("名称（name）已存在，请重新输入莫重复！！！")
			break
		}
	}
	UsersDb[u.Name] = u

	fmt.Printf("用户【%v】添+成功！！！\n", u.Name)
}

func (u *Users) FindUserById(id int) *Users {
	for _, user := range UsersDb {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func (u *Users) ModifyUserById() {
	for _, tuser := range UsersDb {
		if tuser.ID == u.ID {
			if tuser.Name == u.Name {
				fmt.Println("输入的名称已经存在，请重新输入")
				break
			} else {
				delete(UsersDb, tuser.Name)
				UsersDb[u.Name] = u
				fmt.Println("用户修改完成。")
				u.PrintUser()
			}
		}
	}
}

func (u *Users) DeleteUserById() {
	delete(UsersDb, u.Name)
}

func filter(user *Users, q string) bool {
	return strings.Contains(user.Name, q) ||
		strings.Contains(user.Addr, q) ||
		strings.Contains(user.Tel, q)
}

func QueryUser(q string) map[string]*Users {
	rt := make(map[string]*Users)
	for _, user := range UsersDb {
		// fmt.Println(user)
		if filter(user, q) {
			rt[user.Name] = user
		}
	}
	return rt
}

func PrintUsersDb() {
	fmt.Println("\n以下是目前已存在的用户：")
	for _, user := range UsersDb {
		fmt.Printf("ID：%v, 名称：%v, 联系方式：%v, 通信地址：%v，生日：%v\n", user.ID, user.Name, user.Tel, user.Addr, user.Birthday)
	}
}

func (u *Users) PrintUser() {
	fmt.Printf("ID：%v, 名称：%v, 联系方式：%v, 通信地址：%v，生日：%v\n", u.ID, u.Name, u.Tel, u.Addr, u.Birthday)
}
