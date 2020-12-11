package utils

import (
	"flag"
)

//命令行参数解析
func FlagMsg() string {
	//定义命令行参数方式1
	var (
		init string
	)
	flag.StringVar(&init, "i", "csv", "gob/csv/json")

	//解析命令行参数
	flag.Parse()

	// if init == "" {
	// 	flag.Usage()
	// 	// os.Exit(1)
	// }
	dbType := init
	return dbType
}
