package router

import (
	"zhao/controller"
	"zhao/manager"
)

func init() {
	manager.Routers["add"] = controller.AddRun
	manager.Routers["del"] = controller.DelRun
	manager.Routers["modify"] = controller.ModifyRun
	manager.Routers["query"] = controller.QueryRun
	manager.Routers["print"] = controller.PrintAll
	manager.Routers["help"] = controller.View
}
