package routers

import (
	"CMS/controllers"
	"CMS/manager"
)

func init() {
	slice := make([]string, 0, 10)
	manager.Registers(append(slice, "all", "1"), controllers.PringUsersAll)
	manager.Registers(append(slice, "add", "a", "2"), controllers.AddUser)
	manager.Registers(append(slice, "modify", "m", "3"), controllers.ModifyUser)
	manager.Registers(append(slice, "del", "d", "4"), controllers.DeleteUser)
	manager.Registers(append(slice, "query", "q", "5"), controllers.QueryUser)
	manager.Registers(append(slice, "help", "h", "6"), controllers.PringHelpMsg)
	manager.Registers(append(slice, "quit"), controllers.QuitServer)
}
