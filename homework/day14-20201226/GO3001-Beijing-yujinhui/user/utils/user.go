package utils

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"user/models"
)

func GetUserByID(id int64) (*models.User, error) {
	ormer := orm.NewOrm()
	user := models.User{ID: id}
	if err := ormer.Read(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByName(name string) (*models.User, error) {
	ormer := orm.NewOrm()
	user := models.User{Name: name}
	if err := ormer.Read(&user, "Name"); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers() ([]*models.User, error) {
	ormer := orm.NewOrm()
	queryset := ormer.QueryTable(&models.User{})
	var users []*models.User
	_, err := queryset.OrderBy("sort_id").All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func AddUser(name string, password string, addr string, sex bool, sort_id int64) error {

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	user := &models.User{
		Name:     name,
		Password: string(hashed),
		Addr:     addr,
		Sex:      sex,
		Tel:      "1",
		SortID:   sort_id,
	}
	ormer := orm.NewOrm()
	if _, err := ormer.Insert(user); err != nil {
		return err
	}
	return nil
}

func DeleteUserByID(id int64) error {
	if id == 1 {
		return fmt.Errorf("ID为1的用户无法删除")
	}

	ormer := orm.NewOrm()
	if _, err := ormer.Delete(&models.User{ID: id}); err != nil {
		return err
	}
	return nil
}