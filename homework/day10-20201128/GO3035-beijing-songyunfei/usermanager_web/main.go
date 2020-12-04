package main

import (
	"fmt"
	"net/http"
	"path"
	"usermanager_web/controller"
	"usermanager_web/users"
)


func main() {
	var db users.JsonUserDb
	users.Savepath = path.Join("./", "user.json")
	controller.Udb = &db
	addr := "0.0.0.0:8888"
	controller.Templatedir = "./template"
	fmt.Println("启动完成.")

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/add", controller.Add)
	http.HandleFunc("/del", controller.DelUser)
	http.HandleFunc("/modify", controller.ModifyUser)
	http.HandleFunc("/query", controller.Queryuser)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}

}
