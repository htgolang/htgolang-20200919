package main

import (
	"beego_fileserver/config"
	"beego_fileserver/loganalysis"
	"beego_fileserver/process"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	config.InitConfig()
	dbtype := config.Dbtype
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		config.DbUser,config.DbPasswd,config.DbIp,config.DbProt,config.DbName)
	err := process.InitDB(dbtype,dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := process.CloseDb()
		if err != nil {
			fmt.Println(err)
		}
	}()
	addr := fmt.Sprintf("%s:%s",config.Listenaddr,config.Prot)
	fmt.Println("启动完成")
	beego.AutoRouter(&loganalysis.LogAnalysis{})
	beego.Run(addr)

}
