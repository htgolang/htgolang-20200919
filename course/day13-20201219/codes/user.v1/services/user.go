package services

import (
	"fmt"
	"user/config"
	"user/forms"
	"user/models"

	"golang.org/x/crypto/bcrypt"
)

func GetUsers() []*models.User {
	rows, err := config.Db.Query("select id, name, sex, addr from user2")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	users := make([]*models.User, 0, 10)
	for rows.Next() {
		var (
			id   int64
			name string
			sex  bool
			addr string
		)
		if err := rows.Scan(&id, &name, &sex, &addr); err != nil {
			fmt.Println(err)
			break
		}
		users = append(users, models.NewUser(id, name, sex, addr))

	}
	return users
}

func AddUser(name string, addr string, sex bool) {
	result, err := config.Db.Exec("insert into user2(name, sex, addr, created_at, updated_at) values(?, ?, ?, now(), now())", name, sex, addr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}

func DeleteUser(id int64) {
	result, err := config.Db.Exec("delete from user2 where id=?", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}

func GetUserByName(name string) *models.User {
	var (
		uid       int64
		uname     string
		upassword string
	)
	err := config.Db.QueryRow("SELECT id,name,password from user2 where name=?", name).Scan(&uid, &uname, &upassword)
	if err != nil {
		return nil
	}

	user := models.NewUser(uid, uname, false, "")
	user.Password = upassword
	return user
}

func Auth(form *forms.LoginForm) *models.User {
	// 通过用户名去查询用户信息(至少包含密码hash值)
	if user := GetUserByName(form.Username); user == nil {
		// 用户名密码错误
		return nil
	} else {
		// 检查密码
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.Password)); err == nil {
			return user
		} else {
			// 用户名密码错误
			return nil
		}
	}
}
