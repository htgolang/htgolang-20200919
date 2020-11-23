package routers

import (
	"GO3004-zhaoweiping/controllers"
	"GO3004-zhaoweiping/manager"
)

func init() {
	manager.Register("add", controllers.AddUser)
	manager.Register("modify", controllers.ModifyUser)
	manager.Register("del", controllers.DeleteUser)
	manager.Register("query", controllers.QueryUser)
	manager.Register("all", controllers.PringUsersAll)
	manager.Register("help", controllers.PringHelpMsg)
}
