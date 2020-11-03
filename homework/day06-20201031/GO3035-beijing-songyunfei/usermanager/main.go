package main

import (
	"usermanager/controller"
	"usermanager/users"
)

var udb users.Userdb
func init()  {
	udb.Load()
	if len(udb.UserSlice) == 0{
		_ = udb.Add("admin","pek","110","admin","2020-10-29")
		_ = udb.Add("ff","pek","110","ff","2020-10-28")
		udb.Sync()
	}

}

func main()  {
	controller.Run(&udb)
}
