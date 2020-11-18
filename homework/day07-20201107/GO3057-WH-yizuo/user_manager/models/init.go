package models

import (
	"time"
)

/*
初始化结构体及用户数据
*/

type User struct {
	Id       int    // 用户ID
	Name     string // 名称
	Password string // 密码
	Phone    string // 联系方式
	Address  string // 地址
	Birthday time.Time
}

type Users struct {
	curId  int // 唯一
	Status int // 用户状态。0为软删除，1为在使用
	UserData *User//用户列表
}

var UserList = make([]Users ,0)
var UserDataFile string = "models/csv/userData.csv"

func init() {
	/*
	   初始化数据
	   初始化Users，并添加三条基础数据
	   初始化UserPasswd，并转换为十六进制MD5值
	*/
	ReadUsers()
}
