package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	logFilePath = beego.AppConfig.String("log::logFilePath")
)

var logConf = make(map[string]interface{})

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
	log := logs.NewLogger(2048)
	log.SetLevel(logLevel)
	log.SetLogger("file", string(jsonConf))
	log.Info("first log text" + time.Now().String())
}
