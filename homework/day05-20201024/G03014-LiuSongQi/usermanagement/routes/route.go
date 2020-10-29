package routes

import (
	"usermanagement/controllers"
	"usermanagement/manager"
)

func init() {
	manager.Register("add", controllers.AddUser)
	manager.Register("modify", controllers.ModifyUser)
	manager.Register("delete", controllers.DeleteUserById)
	manager.Register("query", controllers.QueryUser)
}
