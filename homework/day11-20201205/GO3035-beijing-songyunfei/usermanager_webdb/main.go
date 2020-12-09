package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"usermanager_webdb/config"
	"usermanager_webdb/controller"
	"usermanager_webdb/users"
)


func main() {
	config.InitConfig()
	mydb := new(users.Userinfo)
	dbtype := config.Dbtype
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
		config.DbUser,config.DbPasswd,config.DbIp,config.DbProt,config.DbName)
	err := mydb.InitDB(dbtype, dsn)

	//log
	logFile, err := os.OpenFile(config.Logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Ldate|log.Ltime)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		err := mydb.CloseDb()
		if err != nil {
			fmt.Println(err)
		}

	}()
	addr := fmt.Sprintf("%s:%s",config.Listenaddr,config.Prot)
	controller.Udb = mydb
	controller.Templatedir = "./template"
	log.Println("启动完成...")

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/add", controller.Add)
	http.HandleFunc("/del", controller.DelUser)
	http.HandleFunc("/modify", controller.ModifyUser)
	http.HandleFunc("/query", controller.Queryuser)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}

}
