package utils

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitLogger() {
	logConf := fmt.Sprintf(`{"filename":"%s","level":%s,"maxlines":%s,"maxsize":%s,"daily":%s,"maxdays":%s,"color":%s}`,
		beego.AppConfig.String("log::log_path"),
		beego.AppConfig.String("log::log_level"),
		beego.AppConfig.String("log::maxlines"),
		beego.AppConfig.String("log::maxsize"),
		beego.AppConfig.String("log::daily"),
		beego.AppConfig.String("log::maxdays"),
		beego.AppConfig.String("log::color"),
	)
	// fmt.Println(logConf)
	// err := logs.SetLogger(logs.AdapterFile, `{"filename":%s,"level":%s,"maxlines":%s,"maxsize":%s,"daily":%s,"maxdays":%s,"color":%s}`)
	err := logs.SetLogger(logs.AdapterFile, logConf)
	if err != nil {
		panic(err)
	}
	logs.Info("CMS 系统启动成功！")
}
