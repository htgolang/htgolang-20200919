package main

import (
	_ "user/database"
	_ "user/router"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
