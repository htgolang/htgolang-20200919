package utils

import (
	"flag"
	"fmt"
	"os"
)

func Flag() {
	var init bool
	flag.BoolVar(&init, "init", false, "初始化")
	flag.Parse()
	if init {
		Load()
		fmt.Println(`初始化程序...`)
		InitPersist()
		InitAdminUser()
		persist := GetPersist()
		persist.Save()
		os.Exit(0)
	}

}