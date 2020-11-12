package utils

import (
	"github.com/olekukonko/tablewriter"
	"sort"
	"strconv"
	"strings"
	"usermanage/model"
)
//定义用户结构体切片类型,以及创建用户信息变量
type UserList []model.User
var UsersList UserList
//定义方法满足sort接口
func (this UserList) Len() int {
	return len(this)
}

func (this UserList) Less(i, j int) bool {
	return this[i].Id < this[j].Id
}

func (this UserList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

//定义用户结构体切片打印
func (this UserList) String() string {
	//对用户列表按Id排序
	sort.Sort(this)
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Id", "Name", "Addr", "Tel", "Birthday", "Password"})
	for _, v := range this {
		table.Append([]string{strconv.Itoa(v.Id), v.Name, v.Addr, v.Tel,
			v.Birthday.Format("2006-01-02"), v.Password})
	}
	table.Render()
	return tableString.String()
}
//初始化admin账户
func InitAdmin() {
	//判断admin账户是否存在，存在则先删除再创建
	//删除之前admin账户
	DelUser("admin")
	passwd := SetPasswd("admin")
	adminUser := model.User{
		Id: 0,
		Name: "admin",
		Password: passwd,
	}
	UsersList = append(UsersList, adminUser)
}
//定义函数map
var FuncMap = map[string]func() {
	"add" : Add,
	"del" : Del,
	"upd" : Update,
	"help" : Help,
	"query" : Query,
	"list" : List,
	"exit" : Exit,
}