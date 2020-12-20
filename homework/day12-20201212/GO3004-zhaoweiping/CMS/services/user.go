package services

import (
	"CMS/models"
	"fmt"
)

func GetUsers() []*models.User {
	rows, err := models.Db.Query("select id, name, sex, addr from user")
	// fmt.Println(rows)
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
		// fmt.Println(birthday.Format("2006-01-02"))
		users = append(users, models.NewUser(id, name, sex, addr))

	}
	// fmt.Println("GetUsers", users)
	return users
}

func AddUser(name string, sex bool, addr string) {
	result, err := models.Db.Exec("insert into user(name, sex, addr, created_at, updated_at) values(?, ?, ?, now(), now())", name, sex, addr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}

func DeleteUser(id int64) {
	result, err := models.Db.Exec("delete from user where id=?", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}

func ModifyUser(id int64, name string, sex bool, addr string) {
	fmt.Println(id, name, sex, addr)
	result, err := models.Db.Exec("update user set name=?, sex=?, addr=?, updated_at=now() where id=?", name, sex, addr, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.LastInsertId())
		fmt.Println(result.RowsAffected())
	}
}
