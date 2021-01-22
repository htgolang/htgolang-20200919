package server

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"zhao/config"
	"zhao/models"

	"golang.org/x/crypto/bcrypt"
)

func GetUsers() []*models.User {
	rows, err := config.DB.Query("select id, name, sex, addr, tel, brithday, passwd, create_at from user")
	if err != nil {
		fmt.Println("[server.GetUsers.Query]", err) //如果数据库出错返回一个空用户列表
		return nil
	}
	defer rows.Close()
	users := make([]*models.User, 0, 20)
	for rows.Next() {
		var (
			id         int64
			name       string
			sex        bool
			addr       string
			tel        string
			brithday   time.Time
			passwd     string
			createTime *time.Time
		)
		err := rows.Scan(&id, &name, &sex, &addr, &tel, &brithday, &passwd, &createTime)
		if err != nil {
			fmt.Println("[server.GetUSers.Scan]", err)
			return nil
		}

		users = append(users, &models.User{
			ID:        id,
			Name:      name,
			Sex:       sex,
			Tel:       tel,
			Addr:      addr,
			Brithday:  brithday.Format("2006-01-02"),
			Passwd:    passwd,
			Create_at: createTime,
		})
	}

	return users
}

func AddUser(name string, sex bool, addr string, tel string, brithday string, passwd string) error {
	bri, err := time.Parse("2006-01-02", brithday)
	if err != nil {
		fmt.Println(err)
		return err
	}
	result, err := config.DB.Exec(
		"insert into user(name, sex, addr, tel, brithday, passwd, create_at) values(?,?,?,?,?,?,now())",
		name, sex, addr, tel, bri, passwd,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if n, err := result.RowsAffected(); n != 1 || err != nil {
		return fmt.Errorf("多行数据")
	}
	return nil
}

func DeleteUser(id int64) error {
	result, err := config.DB.Exec("delete from user where id=?", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n != 1 {
		return fmt.Errorf("删除了%d", n)
	}
	return nil
}

func ParseUser(r *http.Request) (*models.User, error) {
	fmt.Println("ParseUser函数：", r.FormValue("brithday"))
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	_, err := time.Parse("2006-01-02", r.FormValue("brithday"))
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:       id,
		Name:     r.FormValue("name"),
		Sex:      r.FormValue("sex") == "true",
		Tel:      r.FormValue("tel"),
		Brithday: r.FormValue("brithday"),
		Addr:     r.FormValue("addr"),
		Passwd:   r.FormValue("passwd"),
	}, nil
}

func ModifyAuth(user *models.User) error {
	rows, err := config.DB.Query("select name from user")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return err
		}
		if name == user.Name {
			return errors.New("用户名重复")
		}
	}

	return nil
}

func ModifyUser(user *models.User) error {
	result, err := config.DB.Exec("update user set name=?, passwd=?, sex=?, addr=?, tel=?, brithday=? where id=?", user.Name, user.Passwd, user.Sex, user.Addr, user.Tel, user.Brithday, user.ID)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	// fmt.Println("=========>", n)
	if n != 1 {
		return fmt.Errorf("修改%d行的值", n)
	}
	return nil
}

func ParseLogin(basic string) (string, error) {
	up := strings.Fields(basic)[1]
	b := bytes.NewReader([]byte(up))
	reader := base64.NewDecoder(base64.RawStdEncoding, b)
	upb := make([]byte, 1000)
	n, err := reader.Read(upb)
	if err != nil {
		return "", err
	}
	return string(upb[:n]), nil
}

func AuthLogin(userpassword string) bool {
	userpasswd, err := ParseLogin(userpassword)
	if err != nil {
		return false
	}

	i := strings.Index(userpasswd, ":")
	if i == -1 {
		return false
	}
	file, err := os.Open("config/user_passwd")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		confuserpasswd := scanner.Text()
		configi := strings.Index(confuserpasswd, ":")
		if i == -1 {
			continue
		}

		if userpasswd[:i] == confuserpasswd[:configi] {
			if err := bcrypt.CompareHashAndPassword([]byte(strings.TrimSpace(confuserpasswd[configi+1:])), []byte(userpasswd[i+1:])); err != nil {
				return false
			} else {
				return true
			}
		}
	}
	return false
}
