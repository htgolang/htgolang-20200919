package logger

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	logFilePath = beego.AppConfig.String("log::logFilePath")
)

var logConf = make(map[string]interface{})
var Logger = &logs.BeeLogger{}

func InitLogger() {
	logLevel, err := beego.AppConfig.Int("log::logLevel")
	if err != nil {
		fmt.Println(err)
		return
	}
	logConf["filename"] = logFilePath
	jsonConf, errm := json.Marshal(logConf)
	if errm != nil {
		fmt.Println(err)
		return
	}
	Logger = logs.NewLogger(2048)
	Logger.SetLevel(logLevel)
	Logger.EnableFuncCallDepth(true)
	Logger.SetLogger("file", string(jsonConf))
	//Logger.Info("first log text" + time.Now().String())
}
