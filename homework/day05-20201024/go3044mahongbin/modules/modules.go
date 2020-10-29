package modules

import (
	"fmt"
	"go3044/utils"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

//SliceU 保存全部用户的切片
var SliceU []Users

//SliceUID 保存全部用户ID的切片
var SliceUID []int

//Users 定义用户结构体
type Users struct {
	ID                        int
	Name, Password, Tel, Addr string
	Birthday                  time.Time
	Deleted, Ifadmin          bool
}

//AddUser 创建新用户
func AddUser() int {
	var userX Users
	userX.ID = utils.GenID(SliceUID)
START:
	userX.Name = utils.Input("请输入一个新用户名:")
	if IfNameExists(userX.Name) {
		fmt.Println("用户名已存在~")
		goto START
	} else {
		userX.Password = utils.PasswordEncrypt(utils.Input("请输入新用户的密码:"))
		// userX.Password = utils.Input("请输入新用户的密码:")
		userX.Tel = utils.Input("请输入新用户的手机号:")
		userX.Addr = utils.Input("请输入新用户的地址:")
		userX.Birthday = time.Now()
		userX.Deleted = false
		userX.Ifadmin = false
		// 切片变更
		SliceU = append(SliceU, userX)
		SliceUID = append(SliceUID, userX.ID)
		fmt.Println("添加成功!")
		return -1
	}
}

//InitAdmin ...
func InitAdmin() {
	var u Users
	u.ID = utils.GenID(SliceUID)
	u.Name = "adm"
	u.Password = utils.PasswordEncrypt("123123") // admin密码
	// u.Password = "123123"
	u.Tel = "123456789"
	u.Addr = "北京市"
	u.Birthday = time.Now()
	u.Deleted = false
	u.Ifadmin = true // admin账号特权不能被删除
	// 切片变更
	SliceU = append(SliceU, u)
	SliceUID = append(SliceUID, u.ID)
}

//DelUser ...
func DelUser() int {
START:
	// fmt.Println("SliceUID===", SliceUID)
	tmpid, _ := strconv.Atoi(utils.Input("请输入要删除的用户ID:"))
	if tmpid == SliceUID[0] {
		fmt.Println("Admin账号不能被删除!")
		goto START
	}
	// Users是值类型,u是值类型变量
	for i, u := range SliceU {

		if u.ID == tmpid {
			if u.Deleted {
				fmt.Println("该用户已删除,不支持重复删除！")
				return -1
			}
			u.Deleted = true //伪删除
			SliceU[i] = u    // 将修改后的u替换SliceU[i]
			fmt.Println("删除成功！")
			return -1
		}
	}
	fmt.Println("未找到此ID的用户。")
	return -1
}

//ModifyUser ...
func ModifyUser() int {
START:
	// fmt.Println("SliceUID===", SliceUID)
	tmpid, _ := strconv.Atoi(utils.Input("请输入要修改的用户ID:"))
	if tmpid == SliceUID[0] {
		fmt.Println("Admin账号不支持修改!")
		goto START
	}
	found := false
	//值类型
	for idx, u := range SliceU {
		if u.ID == tmpid {
			if u.Deleted {
				fmt.Println("该用户已删除,不支持修改操作！")
				return -1
			} else {
			AGAIN:
				found = true
				tmpF := utils.Input("请在以下用户字段选择一项做修改\nname tel addr\n请输入:")
				switch tmpF {
				case "id":
					fmt.Println("该字段为自动生成，不支持修改！")
					goto AGAIN
				case "birthday":
					fmt.Println("该字段为自动生成，不支持修改！")
					goto AGAIN
				case "name":
					u.Name = utils.Input("请输入一个新名字:")
					if IfNameExists(u.Name) {
						fmt.Println("用户名已存在~")
						goto AGAIN
					} else {
						SliceU[idx] = u // 将修改后的u替换SliceU[i]
						fmt.Println("修改成功！")
						return -1
					}
				case "tel":
					u.Tel = utils.Input("请输入新号码:")
					SliceU[idx] = u // 将修改后的u替换SliceU[i]
					fmt.Println("修改成功！")
					return -1
				case "addr":
					u.Addr = utils.Input("请输入新地址:")
					SliceU[idx] = u // 将修改后的u替换SliceU[i]
					fmt.Println("修改成功！")
					return -1
				default:
					fmt.Println("未找到该字段。。。")
					goto AGAIN
				}
			}
		}
	}
	if !found {
		fmt.Println("未找到此ID的用户。")
	}
	return -1
}

//FindUser ...
func FindUser() int {
	keyword := utils.Input("请输入查询关键字:")
	var result []Users
	for _, u := range SliceU {
		if u.Deleted {
			fmt.Println("该用户已删除！")
			break
		} else {
			switch {
			case strings.Contains(u.Name, keyword):
				result = append(result, u)
			case strings.Contains(u.Tel, keyword):
				result = append(result, u)
			case strings.Contains(u.Addr, keyword):
				result = append(result, u)
			case strings.Contains(u.Birthday.Format("2006-01-02 15:04:05"), keyword):
				result = append(result, u)
			}
		}
	}
	if len(result) == 0 {
		fmt.Println("未找到相关用户。")
	} else {
		tablePrint(result)
	}

	return -1
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

func tablePrint(su []Users) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetAutoFormatHeaders(false)
	t.SetAutoWrapText(false)
	t.SetReflowDuringAutoWrap(false)

	// t.SetHeader([]string{"ID", "Name", "Tel", "Addr", "Deleted", "Birthday"})
	t.SetHeader([]string{"ID", "Name", "Tel", "Addr", "Birthday"})
	for _, u := range su {
		if !u.Deleted {
			// t.Append([]string{strconv.Itoa(u.ID), u.Name, u.Tel, u.Addr, strconv.FormatBool(u.Deleted), u.Birthday.Format("2006-01-02 15:04:05")})
			t.Append([]string{strconv.Itoa(u.ID), u.Name, u.Tel, u.Addr, u.Birthday.Format("2006-01-02 15:04:05")})
		}
	}
	t.Render()
}

//ListUser ...
func ListUser() int {
	tablePrint(SliceU)
	return -1
}

//HelpUser ...
func HelpUser() int {
	fmt.Println(`
	可选操作:
	list	打印用户列表
	add	新增用户
	delete	删除用户
	modify	修改用户
	find	关键字查找
	logout	登出
	exit	退出
	help	帮助
	`)
	return -1
}

//ExitUser ...
func ExitUser() int {
	fmt.Println("再见!")
	return 1
}

//LogoutUser ...
func LogoutUser() int {
	fmt.Println("退出登录!")
	return 0
}
