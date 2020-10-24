package routers

import (
	"github.com/imsilence/user/controllers"
	"github.com/imsilence/user/manager"
)

func init() {
	manager.Register("add", controllers.AddUser)
	manager.Register("modify", controllers.ModifyUser)
	manager.Register("delete", controllers.DeleteUser)
	manager.Register("query", controllers.QueryUser)
	manager.Register("qmc", controllers.QueryMachine)
}
