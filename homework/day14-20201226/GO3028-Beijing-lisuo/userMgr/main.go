package main

import (
"fmt"
	"userMgr/logger"
	_ "userMgr/models"
	_ "userMgr/routers"
	"userMgr/services"

	beego "github.com/astaxie/beego"
)

func main() {
	logger.InitLogger()
	fmt.Printf("dsn: %#v\n", beego.AppConfig.String("db::dsn"))
	services.IfAdmin()
	logger.Logger.Debug("App running...")
	beego.Run()
}
