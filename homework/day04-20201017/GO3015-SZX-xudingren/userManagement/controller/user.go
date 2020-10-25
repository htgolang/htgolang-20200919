package controller

import (
	"fmt"
	"os"
	"strconv"
	"userManagement/utils"

	"github.com/olekukonko/tablewriter"

	"userManagement/model"
	"userManagement/service"
)

type UserView struct {
	arg         string
	userService *service.UserService
}

func NewUserView(service *service.UserService) *UserView {
	return &UserView{
		arg:         "",
		userService: service,
	}
}

func (v *UserView) add() {
	fmt.Println("添加用户")
	fmt.Print("用户名：")
	var name string
	fmt.Scanln(&name)
	fmt.Print("电话：")
	var phone string
	fmt.Scanln(&phone)
	fmt.Print("地址：")
	var address string
	fmt.Scanln(&address)
	user := model.NewUser(name, phone, address)
	if v.userService.Add(user) {
		fmt.Println("添加成功")
	} else {
		fmt.Println("添加失败")
	}
}

func (v *UserView) delete() {
	if v.userService.GetUserNum() == 0 {
		fmt.Println("用户管理列表为空")
		return
	}
	fmt.Println("删除用户")
	inputUserId := utils.Input("请输入关键字：", v.arg)
	uid, _ := strconv.Atoi(inputUserId)
	idx, err := v.userService.HasUser(uid)
	if err != nil {
		fmt.Printf("删除失败：%s\n", err.Error())
		return
	}
	user := v.userService.GetUser(idx)
	v.tableFormat(user)
	cfm := utils.Input("确认删除（y/n）：", v.arg)
	if cfm == "y" {
		if v.userService.Delete(idx) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("删除失败")
		}
	} else {
		fmt.Println("不删除")
	}
}

func (v *UserView) modify() {
	if v.userService.GetUserNum() == 0 {
		fmt.Println("用户管理列表为空")
		return
	}
	fmt.Println("修改用户")
	inputUserId := utils.Input("请输入用户ID：", v.arg)
	uid, _ := strconv.Atoi(inputUserId)
	idx, err := v.userService.HasUser(uid)
	if err != nil {
		fmt.Printf("修改失败：%s\n", err.Error())
		return
	}
	user := v.userService.GetUser(idx)
	v.tableFormat(user)
	cfm := utils.Input("确认修改（y/n）：", v.arg)
	if cfm == "y" {
		name := utils.Input("用户名：", v.arg)
		phone := utils.Input("电话：", v.arg)
		addr := utils.Input("地址：", v.arg)
		mUser := model.User{
			Name:    name,
			Phone:   phone,
			Address: addr,
		}
		if v.userService.Modify(idx, mUser) {
			fmt.Println("修改成功")
		} else {
			fmt.Println("修改失败")
		}
	} else {
		fmt.Println("不修改")
	}
}

func (v *UserView) query() {
	if v.userService.GetUserNum() == 0 {
		fmt.Println("用户管理列表为空")
		return
	}
	fmt.Println("搜索用户")
    keyword := utils.Input("请输入关键字：", v.arg)
	matchUsers := v.userService.Query(keyword)
	if len(matchUsers) != 0 {
		v.tableFormat(matchUsers)
	} else {
		fmt.Println("无匹配用户")
	}
}

func (v *UserView) list() {
	users := v.userService.List()
	v.tableFormat(users)
}

func (v *UserView) menu() {
	fmt.Println("*********用户管理系统*********")
	fmt.Printf("%15s\n", "a）添加用户")
	fmt.Printf("%15s\n", "m）修改用户")
	fmt.Printf("%15s\n", "d）删除用户")
	fmt.Printf("%15s\n", "l）用户列表")
	fmt.Printf("%15s\n", "q）搜索用户")
	fmt.Printf("%15s\n", "h）帮助信息")
	fmt.Printf("%15s\n", "exit）退出系统")
	fmt.Println("****************************")
}

func (v *UserView) tableFormat(data []model.User) {
	fmtData := [][]string{}
	for _, v := range data {
		fmtData = append(fmtData, []string{strconv.Itoa(v.Id), v.Name, v.Phone, v.Address})
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"用户编号", "用户名", "电话号码", "联系地址"})
	for _, v := range fmtData {
		table.Append(v)
	}
	table.Render()
}

func (v *UserView) MainMenu() {
	v.menu()
	for {
		switch utils.Input("输入菜单选项：", v.arg) {
		case "a":
			v.add()
		case "m":
			v.modify()
		case "d":
			v.delete()
		case "l":
			v.list()
		case "q":
			v.query()
		case "h":
			v.menu()
		case "exit":
			fmt.Println("退出用户管理")
			return
		default:
			fmt.Println("非法输入")
		}
	}
}
