package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"usermanager_beego/config"
	"usermanager_beego/controller"
	"usermanager_beego/users"
)


func main() {
	config.InitConfig()
	mydb := new(users.Userinfo)
	dbtype := config.Dbtype
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		config.DbUser,config.DbPasswd,config.DbIp,config.DbProt,config.DbName)

	//初始化日志
	logconf := make(map[string]interface{})
	logconf["filename"] = config.Logfile
	logconf["level"] = config.Log_level
	jsonConfig,err := json.Marshal(logconf)
	if err != nil {
		fmt.Println(err)
		return
	}
	log := logs.NewLogger(1024)
	if err := log.SetLogger(logs.AdapterFile,string(jsonConfig)); err != nil {
		fmt.Println(err)
		return
	}
	log.SetLevel(config.Log_level)
	log.EnableFuncCallDepth(true)
	err = mydb.InitDB(dbtype, dsn)
	if err != nil {
		log.Error("Init DB error:",err)
		return
	}

	addr := fmt.Sprintf("%s:%s",config.Listenaddr,config.Prot)
	controller.Udb = mydb
	log.Info("启动完成...")
	beego.ErrorController(&controller.ErrorController{})
	beego.AutoRouter(&controller.Usermanager{})
	beego.AutoRouter(&controller.Authcontroller{})
	beego.AutoRouter(&controller.LogAnalysis{})
	beego.Router("/",&controller.Usermanager{},"*:Entrance" )
	beego.Run(addr)

}
