package routers

import (
	"yizuo/controllers"
	"yizuo/manager"
)

// 路由代码
func init() {
	/*
		所有用户输入的信息，通过此函数路由至对应接口函数
	*/
	manager.Register("a", controllers.AddUser)
	manager.Register("add", controllers.AddUser)
	manager.Register("d", controllers.DeleteUser)
	manager.Register("delete", controllers.DeleteUser)
	manager.Register("m", controllers.ModifyUser)
	manager.Register("modify", controllers.ModifyUser)
	manager.Register("l", controllers.ListUser)
	manager.Register("list", controllers.ListUser)
	manager.Register("query", controllers.QueryUser)
	manager.Register("init", controllers.InitAllUser)
	manager.Register("h", controllers.HelpPrinting)
	manager.Register("help", controllers.HelpPrinting)
}
