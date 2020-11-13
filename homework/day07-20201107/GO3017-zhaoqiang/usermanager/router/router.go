package router

import (
	"zhao/controller"
	"zhao/manager"
)

func init() {
	manager.Routers["add"] = controller.Add
	manager.Routers["del"] = controller.Del
	manager.Routers["modify"] = controller.ModefyUser
	manager.Routers["query"] = controller.Query
	manager.Routers["help"] = controller.View
	manager.Routers["print"] = controller.PrintA

}
