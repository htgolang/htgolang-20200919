package main

import (

    "fmt"
    "syscall"
    "userManagement/controller"
    "userManagement/service"
    "userManagement/utils"
)

var password = "5f4dcc3b5aa765d61d8327deb882cf99"

func auth() bool {
    return utils.Md5text(utils.Input("请输入密码：", "")) == password
}


//初始化服务
//将服务装载到control
//启动主菜单

func main() {
    if !auth() {
        fmt.Println("密码错误，程序退出")
        syscall.Exit(1)
    }
    userService := service.NewUserService()
    userCtrl := controller.NewUserView(userService)
    userCtrl.MainMenu()
}
