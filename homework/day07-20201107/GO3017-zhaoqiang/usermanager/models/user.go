package models

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"zhao/utils"

	"github.com/olekukonko/tablewriter"
)

var users = make([]User, 0, 10)

// User  user infomations
// 	ID string
// 	Name     string
// 	Addr     string
// 	Tel      string
// 	Birthday time.Time
// 	Passwd   string
type User struct {
	ID       int
	Name     string
	Addr     string
	Tel      string
	Birthday time.Time
	Passwd   string
}

// ------
//add用户
func getUserID() int {
	var id int
	for _, user := range users {
		if user.ID > id {
			id = user.ID
		}
	}
	return id + 1
}

//GetUserInfo get user commandline input message
func GetUserInfo(flag string) (User, error) {
	var id int
	if flag == "add" {
		id = getUserID()
	} else if flag == "modify" {
		input := utils.GetInput("input user id: ")
		idint, err := strconv.Atoi(input)
		if err != nil {
			return User{}, fmt.Errorf("id must integr")
		}
		if idint < 0 {
			return User{}, fmt.Errorf("id must greater zone")
		}
		id = idint
	}

	user := User{
		ID:   id,
		Name: utils.GetInput("input username: "),
		Addr: utils.GetInput("input addr: "),
		Tel:  utils.GetInput("input phone: "),
	}
	birthdaystr := utils.GetInput("input brithday(2020年11月12日/20201112):")
	passwd1, err := utils.GetPasswd("input passwd: ")
	passwd2, err := utils.GetPasswd("input agent passwd: ")
	//密码判断
	if passwd1 != passwd2 {
		return User{}, fmt.Errorf("密码不一致")
	}
	if len(passwd1) < 1 {
		return User{}, fmt.Errorf("密码不能为空")
	}
	//生日格式判断
	brithday, err := time.Parse("20060102", birthdaystr)
	if err != nil {
		return User{}, err
	}

	passwd := utils.Md5Convert([]byte(passwd2))
	user.Birthday = brithday
	user.Passwd = passwd
	return user, nil
}

//AddUser adduser action
func AddUser(user User) error {
	users = append(users, user)
	if err := File.Storage(users); err != nil {
		return err
	}
	return nil
}

//AddAuth 添加用户判断id是否有相同
func AddAuth(u User) error {
	for _, user := range users {
		if u.ID == user.ID {
			return fmt.Errorf("user id already exist")
		}
	}
	return nil
}

//GenericQuery 查询
func GenericQuery(text string) []User {
	tempUsers := make([]User, 0, 10)
	for _, user := range users {
		if filter(user, text) {
			tempUsers = append(tempUsers, user)
		}
	}
	return tempUsers
}

func filter(user User, text string) bool {
	id := strconv.Itoa(user.ID)
	birthdaystr := user.Birthday.Format("20060102")
	if id == text ||
		strings.Contains(user.Name, text) ||
		strings.Contains(user.Addr, text) ||
		user.Tel == text ||
		birthdaystr == text {
		return true
	}
	return false
}

//IDQuery query by id
func IDQuery(text string) (User, int, error) {
	id, err := strconv.Atoi(text)
	if err != nil {
		return User{}, -1, err
	}
	for indes, user := range users {
		if id == user.ID {
			return user, indes, nil
		}
	}
	return User{}, -1, fmt.Errorf("no this user")
}

//DelUser deluser action
func DelUser(u User) error {
	tmpUsers := make([]User, 0, len(users)-1)
	for _, user := range users {
		if user.ID != u.ID {
			tmpUsers = append(tmpUsers, user)
		}
	}
	users = tmpUsers

	if err := File.Storage(users); err != nil {
		return err
	}
	return nil
}

//ModifyUser modify user action
func ModifyUser(oldindex int, new User) error {
	err := ModifyUserAuth(oldindex, new)
	if err == nil {
		users[oldindex] = new

		if err := File.Storage(users); err != nil {
			return err
		}
		return nil
	}
	return err
}

//ModifyUserAuth 判断要修改的用户id是否相同
func ModifyUserAuth(oldindex int, new User) error {
	for index, user := range users {
		if index == oldindex {
			continue
		}
		if user.ID == new.ID {
			return fmt.Errorf("user id already exist")
		}
	}
	return nil
}

//ModifypasswdAuther 修改用户前认证权限用户为超级用户，直接修改， 普通用户输入要修改用户的密码
func ModifypasswdAuther(u User) error {
	if LoginUser.Name == "admin" {
		return nil
	} else if u.Name == LoginUser.Name {
		return nil
	} else if u.Name != LoginUser.Name {
		prompt := "enter " + u.Name + " passwd: "
		pass, err := utils.GetPasswd(prompt)
		if err != nil {
			return err
		}
		if utils.Md5Convert([]byte(pass)) == u.Passwd {
			return nil
		}
		return fmt.Errorf("%s passwd miss", u.Name)

	}
	return fmt.Errorf("Unknow err")
}

//PrintUsers 打印用户信息
func PrintUsers(us []User) {
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetAutoFormatHeaders(false)
	tw.SetAutoWrapText(false)
	tw.SetHeader([]string{"编号", "名字", "地址", "电话", "生日", "密码"})
	for _, user := range us {
		brithstr := user.Birthday.Format("2006/01/02")
		tw.Append([]string{strconv.Itoa(user.ID), user.Name, user.Addr, user.Tel, brithstr, user.Passwd})
	}

	tw.Render()
}

//PrintAllusers 打印所有的用户信息
func PrintAllusers() {
	PrintUsers(users)
}
