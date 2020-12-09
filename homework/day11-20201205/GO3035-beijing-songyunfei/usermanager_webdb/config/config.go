package config

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"path"
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

func InitConfig()  {
	//cmdpath,err := os.Executable()
	//if err != nil{
	//	panic("err")
	//}
	//dir,_ := filepath.Split(cmdpath)
	inipath := path.Join("./","config.ini")
	Config,err := goconfig.LoadConfigFile(inipath)
	if err !=nil{
		fmt.Println("Load config.ini  error:",err)
	}
	Dbtype, err = Config.GetValue("DB", "dbtype")
	checkerr(err)
	DbIp, err = Config.GetValue("DB", "ipaddr")
	checkerr(err)
	DbProt, err = Config.GetValue("DB", "dbprot")
	checkerr(err)
	DbName, err = Config.GetValue("DB", "dbname")
	checkerr(err)
	DbUser, err = Config.GetValue("DB", "dbuser")
	checkerr(err)
	DbPasswd, err = Config.GetValue("DB", "dbpasswd")
	checkerr(err)
	Listenaddr, err = Config.GetValue("App", "ListenAddr")
	checkerr(err)
	Prot, err = Config.GetValue("App", "Prot")
	checkerr(err)
	Logfile, err = Config.GetValue("Log", "logfile")
	checkerr(err)


}

func checkerr(err error)  {
	if err !=nil{
		fmt.Println("Error:",err)
	}
}
