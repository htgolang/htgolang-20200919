package services

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
	"usermanage/config"
	"usermanage/model"
)

// 新增用户函数
func AddUser(addpage *model.AddUpdatePage) {
	var sex bool
	if addpage.IsMale == ""{
		sex = false
	} else {
		sex = true
	}
	_, err := config.Db.Exec("insert into user(name,sex,addr,tel,birthday,password) values (?,?,?,?,?,?)",
		addpage.Name, sex, addpage.Addr, addpage.Tel, addpage.Birthday,
		fmt.Sprintf("%x", md5.Sum([]byte(addpage.Passwd))))
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 读取全部用户信息函数
func GetAllUser() []*model.User {
	rows, err := config.Db.Query("select id,name,sex,addr,tel,birthday,password from user")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer rows.Close()
	users := make([]*model.User, 0, 10)
	for rows.Next() {
		var (
			id   int
			name string
			sex  bool
			addr string
			tel string
			birthday time.Time
			password string
		)
		if err := rows.Scan(&id, &name, &sex, &addr, &tel, &birthday, &password); err != nil {
			fmt.Println(err)
			break
		}
		users = append(users, model.NewUser(id, name, sex, addr, tel, birthday, password))

	}
	return users
}

// 根据id查询用户
func GetPageById(id int, page *model.AddUpdatePage)  {
	rows, err := config.Db.Query("select name, sex, addr, tel, birthday, password from user where id=?", id)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	var (
		name string
		sex  bool
		addr string
		tel string
		birthday time.Time
		password string
	)
	for rows.Next() {
		if err := rows.Scan(&name, &sex, &addr, &tel, &birthday, &password); err != nil {
			fmt.Println(err)
			break
		}
		page.Name = name
		if sex {
			page.IsMale = "checked"
			page.IsFeMale = ""
		} else {
			page.IsMale = ""
			page.IsFeMale = "checked"
		}
		page.Tel = tel
		page.Addr = addr
		page.Birthday = birthday.Format("2006-01-02")
		page.Passwd = password
	}
}

// 定义更新用户函数
func UpdateUser(id int, page *model.AddUpdatePage) {
	var sex bool
	if page.IsMale == ""{
		sex = false
	} else {
		sex = true
	}
	sql := fmt.Sprintf("update user set name='%v',sex=%v,addr='%v',tel='%v',birthday='%v',password='%v' where id=?",
		page.Name, sex, page.Addr, page.Tel, page.Birthday,
		fmt.Sprintf("%x", md5.Sum([]byte(page.Passwd))))
	_, err := config.Db.Exec(sql, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 删除用户
func DeleteUser(id int) {
	_, err := config.Db.Exec("delete from user where id=?", id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 条件查询用户

func GetQueryUser(mainpage *model.MainPage) []*model.User {
	filtersql := ""
	filters := []string{}
	if mainpage.Id != "" {
		id, _ := strconv.Atoi(mainpage.Id)
		filters = append(filters, fmt.Sprintf("id=%v", id))
	}
	if mainpage.Name != "" {
		filters = append(filters, fmt.Sprintf("name='%v'", mainpage.Name))
	}
	if mainpage.Sex != "" {
		if mainpage.Sex == "男" {
			filters = append(filters, fmt.Sprintf("sex=1"))
		} else if mainpage.Sex == "女" {
			filters = append(filters, fmt.Sprintf("sex=0"))
		}
	}
	if mainpage.Addr != "" {
		filters = append(filters, fmt.Sprintf("addr='%v'", mainpage.Addr))
	}
	if mainpage.Birthday != "" {
		filters = append(filters, fmt.Sprintf("birthday='%v'", mainpage.Birthday))
	}
	if mainpage.Tel != "" {
		filters = append(filters, fmt.Sprintf("tel='%v'", mainpage.Tel))
	}

	for _, filter := range filters {
		if filtersql == "" {
			filtersql += " where " + filter
		} else {
			filtersql += " and " + filter
		}
	}

	rows, err := config.Db.Query("select id,name,sex,addr,tel,birthday,password from user" + filtersql)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()
	users := make([]*model.User, 0, 10)
	for rows.Next() {
		var (
			id   int
			name string
			sex  bool
			addr string
			tel string
			birthday time.Time
			password string
		)
		if err := rows.Scan(&id, &name, &sex, &addr, &tel, &birthday, &password); err != nil {
			fmt.Println(err)
			break
		}
		users = append(users, model.NewUser(id, name, sex, addr, tel, birthday, password))

	}
	return users
}
