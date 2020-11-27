package main

import (
	"flag"
	"fmt"
	//"io"
	"net"
	"os"
	"path"
	"usermanager_net/controller"
	"usermanager_net/users"
	"usermanager_net/userutils"
)

var udb users.Mydb
func main()  {
	var r = flag.Bool("init",false,"初始化")
	var p = flag.String("p","./","指定用户文件存储目录,默认为当前目录")
	var t = flag.String("t","json","指定用户文件存储类型. 支持json, csv, 默认json.")
	var l = flag.Int("l",3,"指定保存最后n次更改")
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
	default:
		fmt.Printf("不支持的类型:%s\n",*t)
		return
	}
	if *r {
		isexist := func(s string) bool{
			_,err := os.Stat(s)
			if err == nil {
				return true
			}
			if os.IsNotExist(err){
				return false
			}
			return false
		}
		if isexist(users.Savepath){
			y := userutils.Input("DB文件已经存在确认要覆盖吗? [y/n]:")
			if y == "y" || y == "Y" {
				file,err := os.OpenFile(users.Savepath,os.O_TRUNC|os.O_RDWR|os.O_CREATE,os.ModePerm)
				if err !=nil{
					fmt.Println(fmt.Errorf("打开文件失败:%s",err))
				}
				users.Filefd = file
				_ = users.Filefd.Sync()
				//controller.Add(udb)
			}
		}

		file,err := os.OpenFile(users.Savepath,os.O_RDWR|os.O_CREATE,os.ModePerm)
		if err !=nil{
			fmt.Println(fmt.Errorf("打开文件失败:%s",err))
		}
		defer func() {
			_ = file.Close()
		}()
		users.Filefd = file
		err = udb.Load()
		//if  err == io.EOF {
		//	controller.Add(udb)
		//}


	}else {
		file,err := os.OpenFile(users.Savepath,os.O_CREATE|os.O_RDWR,os.ModePerm)
		defer func() {
			_ = file.Close()
		}()
		if err !=nil{
			fmt.Println(fmt.Errorf("打开文件失败:%s",err))
		}
		users.Filefd = file
		err = udb.Load()
		if err != nil {
			fmt.Println(err)
			return
		}

	}
	users.QueueLen = *l
	addr := "0.0.0.0:8888"
	ptc := "tcp"
	lisen,err := net.Listen(ptc,addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("启动完成.")
	for {
		conn,err := lisen.Accept()
		if err !=nil{
			fmt.Println(err)
			return
		}
		fmt.Printf("%s 已连接...\n",conn.RemoteAddr())
		go controller.Run(conn,udb)
	}


}
