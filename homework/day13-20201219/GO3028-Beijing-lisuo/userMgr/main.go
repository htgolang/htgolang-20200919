package main

import (
	"userMgr/logger"
	_ "userMgr/models"
	_ "userMgr/routers"
	"userMgr/services"

	beego "github.com/astaxie/beego"
)

func main() {
	logger.InitLogger()
	//models.Init()
	//var driverName = "mysql"
	//var dsn = "web:web@tcp(127.0.0.1:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"
	//if err := models.InitDB(driverName, dsn); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer models.CloseDB()
	//// test
	//o := orm.NewOrm()
	//fmt.Printf("orm: %#v\n", o)
	//qs := o.QueryTable("user")
	//fmt.Printf("qs: %#v\n", qs)
	//fmt.Println(qs.Count())
	//var users = []*models.User{}
	//n, err := qs.All(&users)
	//fmt.Printf("query result: %#v, err: %#v\n", n, err)
	//for k, v := range users {
	//	fmt.Printf("k: %#v, v: %#v", k, v)
	//}
	////
	//fmt.Printf("dsn: %#v\n", beego.AppConfig.String("db::dsn"))
	services.IfAdmin()

	beego.Run()
}
