package config

import (
	"fmt"
	"github.com/astaxie/beego"
)

var Dbtype string
var DbIp string
var DbProt string
var DbName string
var DbUser string
var DbPasswd string
var Listenaddr string
var Prot string
var Logfile string
var Log_level int

func InitConfig()  {
	Dbtype  = beego.AppConfig.String("DB::dbtype")
	DbIp  = beego.AppConfig.String("DB::ipaddr")
	DbProt = beego.AppConfig.String("DB::dbprot")
	DbName = beego.AppConfig.String("DB::dbname")
	DbUser = beego.AppConfig.String("DB::dbuser")
	DbPasswd = beego.AppConfig.String("DB::dbpasswd")
	Listenaddr = beego.AppConfig.String("App::ListenAddr")
	Prot = beego.AppConfig.String("App::Prot")
	Logfile = beego.AppConfig.String("log::log_path")
	var err error
	Log_level,err = beego.AppConfig.Int("log::log_level")
	if err != nil{
		fmt.Println(err)
		return
	}

}
