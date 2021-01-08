package server

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"user/models"

	"github.com/astaxie/beego/orm"
)

func GetUsers() ([]*models.User, error) {
	ormer := orm.NewOrm()
	queryseter := ormer.QueryTable(new(models.User))
	users := make([]*models.User, 0, 50)
	_, err := queryseter.All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *models.User) error {
	pass, err := models.PassCrypto(user.PassWord)
	if err != nil {
		return err
	}
	user.PassWord = pass
	ormer := orm.NewOrm()
	_, err = ormer.Insert(user)
	if err != nil {
		return err
	}
	return nil

}

func DeleteUserByID(id int64) error {
	user := &models.User{
		ID: id,
	}
	ormer := orm.NewOrm()
	_, err := ormer.Delete(user, "ID")
	if err != nil {
		return err
	}
	return nil
}

func QueryUserByID(i int64) (*models.User, error) {
	ormer := orm.NewOrm()
	user := &models.User{
		ID: i,
	}
	err := ormer.Read(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func QueryUserByName(name string) (*models.User, error) {
	user := &models.User{Name: name}
	ormer := orm.NewOrm()
	err := ormer.Read(user, "Name")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func ModifyUser(user *models.User) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), 5)
	if err != nil {
		return err
	}
	user.PassWord = string(pass)

	ormer := orm.NewOrm()
	_, err = ormer.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func AuthName(user *models.User, action func(user *models.User) error) error {
	users, err := GetUsers()
	if err != nil {
		return err
	}
	for _, u := range users {
		if u.ID == user.ID {
			continue
		}
		if u.Name == user.Name {
			return errors.New("用户名相同")
		}
	}
	err = action(user)
	if err != nil {
		return err
	}
	return nil
}

func AuthLogin(user *models.LoginUser) (int64, error) {
	u, err := QueryUserByName(user.Name)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(user.PassWord))
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

func ModifyUserPerm(user *models.User) error {
	ormer := orm.NewOrm()
	_, err := ormer.Update(user)
	if err != nil {
		return err
	}
	return nil
}
