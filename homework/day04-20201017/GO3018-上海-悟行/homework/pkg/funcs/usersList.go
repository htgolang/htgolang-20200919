package funcs

import (
	"os"
	"homework/pkg/models"
	"github.com/olekukonko/tablewriter"
)
//表格方式输出所有数据
func usersList(Users *[]map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID","Name","Contact","Address"})
	for _,v := range * Users {
		table.Append(models.ConvertElementToSlice(v))
	}
	table.Render()
}
//根据ID返回数据
func userIdList(userID int) {
	data := []map[string]string{models.Users[userID]}
	usersList(&data)
}