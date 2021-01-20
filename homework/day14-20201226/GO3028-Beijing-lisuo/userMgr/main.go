package main

import (
	"fmt"
	"userMgr/logger"
	"userMgr/services"

	"userMgr/models"
	_ "userMgr/routers"

	beego "github.com/astaxie/beego"
	orm "github.com/astaxie/beego/orm"
)

func init() {
    models.RegisterDB()
}

func main() {
    orm.Debug = true
	logger.InitLogger()
	fmt.Printf("dsn: %#v\n", beego.AppConfig.String("db::dsn"))
	services.IfAdmin()
	logger.Logger.Debug("App running...")
	beego.Run()
}
