package main

import (
	_ "userMgr/routers"

	beego "github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
