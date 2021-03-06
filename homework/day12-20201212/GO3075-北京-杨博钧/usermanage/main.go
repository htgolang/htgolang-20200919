package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"usermanage/config"
	"usermanage/routers"
)

func main() {
	// 定义端口绑定5849
	addr := ":5849"

	// 获取数据库配置
	Dbconf := config.NewDbConf()

	// 连接数据库
	err := Dbconf.InitDb()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer config.CloseDb()

	// 注册页面
 	routers.Register()

	// 启动监听服务
	beego.Run(addr)
}

