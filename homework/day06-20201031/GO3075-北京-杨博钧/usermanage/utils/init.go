package utils

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
	"strconv"
	"strings"
	"usermanage/model"
	"io/ioutil"
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
//定义文件持久化接口
type Persist interface {
	Save()
	Load()
}
//定义初始化存储设置
func InitPersist() (Persist) {
	mod := ""
	data, err := ioutil.ReadFile("mod.conf")
	if err != nil {
		fmt.Println("无法打开保存格式配置文件，默认使用csv作为持久化格式")
		mod = "csv"
	}
	mod = string(data)
	var persist Persist
	switch mod {
	case "gob":
		persist = GobPersister{}
	case "csv":
		persist = CsvPersister{}
	default :
		fmt.Printf("无法识别%v持久化格式,请以gob/csv/json作为持久化格式，退出程序\n", mod)
		os.Exit(0)
	}
	return persist
}