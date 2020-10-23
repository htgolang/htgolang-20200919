package funcs

import (
	"GO3057-WH-yizuo/pkg/models"
	"os"

	"github.com/olekukonko/tablewriter"
)

// 将现在所有的用户及其相关数据，格式化输出
func UsersList(Users *[]map[string]string) {
	/*
	   数据格式化相关疑问可以查看  https://github.com/olekukonko/tablewriter
	*/
	table := tablewriter.NewWriter(os.Stdout)
	// 定义表格标题
	table.SetHeader([]string{"ID", "Name", "Contact", "Address"})
	// 遍历数据添加至表格中
	for _, v := range *Users {
		// 添加数据至表格中（现将遍历的字典数据转换为slice切片）
		table.Append(models.ConvertElementToSlice(v))
	}
	// 输出数据
	table.Render()
}

func userIdList(CAP int) {
	/*
		根据用户ID返回该ID对应的数据条目
	*/
	data := []map[string]string{models.Users[CAP]}
	UsersList(&data)
}
