package controller

import (
	"fmt"
	"zhao/models"
)

//AuthenPasswdCount 密码认证次数限制
func AuthenPasswdCount(passwdcount int) (string, bool) {
	for i := 0; i < passwdcount; i++ {
		if useraccount, ok := models.AuthenPassWD(); ok {
			return useraccount, true
		}
		fmt.Printf("Permission denied,  trey again.\n\n")
	}
	return "", false
}
