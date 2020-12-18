package utils

import (
	"fmt"
	"usermanage/config"
	"usermanage/forms"
)

// 判断添加一个用户是否存在
func IsAddUserExists(name string) (bool) {
	rows, err := config.Db.Query("select count(*) from user where name=?", name)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer rows.Close()
	var c int
	for rows.Next() {
		err := rows.Scan(&c)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return c > 0
}

// 判断添加一个更新用户是否存在
func IsUpdateUserExists(id int,name string) (bool) {
	rows, err := config.Db.Query("select count(*) from user where name=? and id!=?", name, id)
	if err != nil {
		fmt.Println(err)
		return true
	}
	defer rows.Close()
	var c int
	for rows.Next() {
		err := rows.Scan(&c)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return c > 0
}

// 判断两次密码是否一致
func IsPasswdSame(user forms.User) (bool) {
	return user.Password == user.Confirm
}
