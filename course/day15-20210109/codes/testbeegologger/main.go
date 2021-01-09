package main

import (
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{
		"level":7,
		"filename":"test.log",
		"maxlines": 80
	}`)
	for i := 0; i < 100; i++ {
		beego.Debug("debug")
		beego.Informational("info")
		beego.Warning("warning")
		beego.Error("error")
	}
}
