package models

import (
	"fmt"
)

func init() {
	fmt.Println("初始化用户！！！请等待")
	// adminBirthdy := utils.TranceTime("2020-11-20 22:40:00")
	// adminPwd := utils.Md5text("1")
	// users1 := NewUser(1, "admin", "adminAddr", "adminTel", adminBirthdy, adminPwd)
	// users1.AddUser()
	// users2 := NewUser(2, "test", "testAddr", "testTel", adminBirthdy, adminPwd)
	// users2.AddUser()
	// JsonToDb()
	// CsvToDb()
	GobToDb()
	// fmt.Println(UsersDb)
}
