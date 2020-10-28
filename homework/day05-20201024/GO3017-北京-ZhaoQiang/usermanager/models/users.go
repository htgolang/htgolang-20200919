package models

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"zhao/utils"

	"github.com/olekukonko/tablewriter"
)

const passWD string = "4e7d489b49ec93dbf53ce37aee778593" //123@qwe

var users = []User{}

// User  user infomations
// 	id string
// 	name     string
// 	addr     string
// 	tel      string
// 	birthday time.Time
// 	passwd   string
type User struct {
	id       int
	name     string
	addr     string
	tel      string
	birthday time.Time
	passwd   string
}

func init() {
	adminBrith, _ := time.Parse("2006/01/02", "2020/10/25")

	admin := User{
		id:       0,
		name:     "admin",
		addr:     "admin",
		tel:      "admin",
		birthday: adminBrith,
		passwd:   passWD,
	}

	users = append(users, admin)
}

//GetUserMessage 获取用户输入的 用户信息
func GetUserMessage() (User, error) {
	id := getID() + 1
	name := utils.GetUserInputString("enter username: ")
	address := utils.GetUserInputString("enter address: ")
	telphone := utils.GetUserInputString("enter telphone: ")
	brithday := utils.GetUserInputTime("enter your brithday(2006/01/02): ")
	passwd1 := utils.GetUserInputPassWD("enter your passwd: ")
	passwd2 := utils.GetUserInputPassWD("reenter passwd: ")
	if passwd1 == passwd2 {
		user := User{
			id:       id,
			name:     name,
			addr:     address,
			tel:      telphone,
			birthday: brithday,
			passwd:   passwd2,
		}
		user.passwd = utils.Md5Text(user.passwd) //修改铭文密码为md5
		return user, nil
	}
	return users[0], errors.New("There is a difference between the two password")
}

//Printfunc 打印函数
func Printfunc(us []User) {
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetAutoFormatHeaders(false)
	tw.SetAutoWrapText(false)

	tw.SetHeader([]string{"编号", "名字", "地址", "电话", "生日", "密码"})

	for _, user := range us {
		brithstr := user.birthday.Format("2006/01/02")
		tw.Append([]string{strconv.Itoa(user.id), user.name, user.addr, user.tel, brithstr, user.passwd})
	}

	tw.Render()
}

// PrintAll print us infomation
func PrintAll() {
	Printfunc(users)
}

//------------------------------------------------------------------------------------------------------------

// Query find all contains user message
func Query(order string) []User {

	tmpusers := make([]User, 0, 10)
	for _, user := range users {
		if filter(user, order) {
			tmpusers = append(tmpusers, user)
		}
	}
	return tmpusers
}

func filter(u User, q string) bool {
	id := strconv.Itoa(u.id)
	brither := u.birthday.Format("2006/01/02")
	if strings.Contains(id, q) ||
		strings.Contains(u.name, q) ||
		strings.Contains(u.tel, q) ||
		strings.Contains(u.addr, q) ||
		strings.Contains(brither, q) {
		return true
	}
	return false
}

//------------------------------------------------------------------------------------------------------------

// ModifyUserMessageByID +++
func ModifyUserMessageByID(olduser, newuser User) error {
	var tmpi int
	for index, user := range users {

		if user == olduser {
			tmpi = index
			continue
		}
		//验证修改用户的id值是否相同
		if user.id == newuser.id {
			return errors.New("用户的id已经存在")
		}
		//用户名是否相同
		if user.name == newuser.name {
			return errors.New("用户名已经存在")
		}
	}

	users[tmpi] = newuser
	return nil
}

//GetModifyUserMessage 获取用户输入的修改信息
func GetModifyUserMessage() (User, error) {
	id := utils.GetUserInputInt("enter user id[modify]: ")

	userinputmessage, err := GetUserMessage()
	if err != nil {
		fmt.Printf("%v\n\n", err)
		return users[0], err
	}
	userinputmessage.id = id
	return userinputmessage, nil

}
