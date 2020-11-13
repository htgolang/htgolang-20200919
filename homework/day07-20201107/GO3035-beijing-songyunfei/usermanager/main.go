package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"usermanager/controller"
	"usermanager/users"
	"usermanager/userutils"
)

var udb users.Mydb
func main()  {
	var r = flag.Bool("init",false,"初始化")
	var p = flag.String("p","./","指定用户文件存储目录,默认为当前目录")
	var t = flag.String("t","json","指定用户文件存储类型. 支持json, csv, 默认json.")
	flag.Parse()
	switch *t {
	case "json":
		var db users.JsonUserDb
		udb = &db
		users.Savepath = path.Join(*p,"user.json")
	case "csv":
		var db users.CsvUserDb
		udb = &db
		users.Savepath = path.Join(*p,"user.csv")
		fmt.Println(users.Savepath)
	default:
		fmt.Printf("不支持的类型:%s\n",*t)
		return
	}
	if *r {
		if err := udb.Load(); err == io.EOF {
			controller.Add(udb)
		}else if err == nil {
			y := userutils.Input("DB文件已经存在确认要覆盖吗? [y/n]:")
			if y == "y" || y == "Y" {
				file,err := os.OpenFile(users.Savepath,os.O_TRUNC|os.O_WRONLY|os.O_CREATE,os.ModePerm)
				if err !=nil{
					fmt.Println(fmt.Errorf("打开文件失败:%s",err))
				}
				_ = file.Sync()
				_ = file.Close()
				controller.Add(udb)
			}
		}
	}else {
		err := udb.Load()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	controller.Run(udb)
}
