package models

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"reflect"
)

var userKay = []string{"ID", "Name","Password", "Phone", "Address","Birthday"}

// User类型的数据格式化输出
func FormatTableOut(u *User) {
	/*
	   数据格式化相关疑问可以查看  https://github.com/olekukonko/tablewriter
	*/
	table := tablewriter.NewWriter(os.Stdout)
	// 定义表格标题
	table.SetHeader(userKay)
	// 添加数据至表格中（现将遍历的字典数据转换为slice切片）
	table.Append(ConvertElementToSlice(u))
	// 输出数据
	table.Render()
}

// Users类型的数据格式化输出
func FormatListTableOut(u []Users) {
	/*
	   数据格式化相关疑问可以查看  https://github.com/olekukonko/tablewriter
	*/
	table := tablewriter.NewWriter(os.Stdout)
	// 定义表格标题
	table.SetHeader(userKay)
	// 遍历数据添加至表格中
	for _, v := range u {
		// 添加数据至表格中（现将遍历的字典数据转换为slice切片）
		table.Append(ConvertElementToSlice(v.UserData))
	}
	// 输出数据
	table.Render()
}

// 结构体转换成切片
func ConvertElementToSlice(element *User) (ret []string) {
	/*
	   结构体类型的数据转换为切片类型
	*/
	ret = make([]string, 0)
	var value = reflect.ValueOf(*element)
	for i := 0; i < value.NumField(); i++ {
		data := fmt.Sprintf("%v",value.Field(i))
		ret = append(ret,data)
	}
	return
}