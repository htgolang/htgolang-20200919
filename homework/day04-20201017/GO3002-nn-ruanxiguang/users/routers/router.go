package routers

import (
	"users/controllers"
	"users/manager"
	"users/models"
)

func init() {
	manager.Register("add", controllers.AddUser)
	manager.Register("delete", controllers.DeleteUser)
	manager.Register("modify", controllers.ModifyUser)
	manager.Register("query", controllers.QueryUser)
	manager.Register("help", models.Help)
}
