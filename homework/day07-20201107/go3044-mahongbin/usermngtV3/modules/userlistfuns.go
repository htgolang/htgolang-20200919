package modules

import (
	"os"
	"strconv"
	"time"
	"usermanage/utils"

	"github.com/olekukonko/tablewriter"
)

//InitAdmin 生成adm用户
func InitAdmin() {
	var u Users
	u.ID = utils.GenID(SliceUID)
	u.Name = "adm"
	u.Password = utils.PasswordEncrypt("123123") // admin密码
	// u.Password = "123123"
	u.Tel = "15166668888"
	u.Addr = "北京市"
	u.Birthday = time.Now()
	u.Deleted = false
	u.Ifadmin = true // admin账号特权不能被删除
	// 切片变更
	SliceU = append(SliceU, u)
	SliceUID = append(SliceUID, u.ID)
}

//UserAuth ...
func UserAuth(tmpN, tmpP string) bool {
	authOK := false
	enP := utils.PasswordEncrypt(tmpP)
	for _, u := range SliceU {
		if u.Name == tmpN && u.Password == enP {
			authOK = true
			break
		}
	}
	return authOK
}

//IfNameExists ...
func IfNameExists(name string) bool {
	for _, su := range SliceU {
		if su.Name == name {
			return true
		}
	}
	return false
}

//TablePrint 以表格形式打印输出
func TablePrint(su []Users) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	// t.SetHeader([]string{"ID", "Name", "Tel", "Addr", "Deleted", "Birthday"})
	t.SetHeader([]string{"ID", "Name", "Tel", "Addr", "Birthday"})
	for _, u := range su {
		if !u.Deleted {
			// t.Append([]string{strconv.Itoa(u.ID), u.Name, u.Tel, u.Addr, u.Birthday.Format("2006-01-02 15:04:05")})
			t.Append([]string{strconv.Itoa(u.ID), u.Name, u.Tel, u.Addr, u.Birthday.Format("2006-01-02 15:04:05")})
		}
	}
	t.Render()
}
