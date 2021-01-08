package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "user/models"
	_ "user/routers"
	"user/utils"
)

func main() {
	// 同步数据库
	orm.RunSyncdb("default", false, true)

	if _, err := utils.GetUserByID(1); err != nil {
		utils.AddUser("admin", "123456", "河北保定", true, 1)
	}
	beego.Run("127.0.0.1:80")
}
