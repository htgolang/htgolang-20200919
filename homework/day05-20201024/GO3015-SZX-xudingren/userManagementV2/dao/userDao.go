package dao

import (
	"userManagementV2/model"
)

//DAO

type UserDao interface {
	Create(name, phone, addr string) error
	Get(userId int) (*model.User, error)
	Update(userId int, name, phone, addr string) error
	List() ([]*model.User, error)
	Query(key string) ([]*model.User, error)
}

type userDaoImp struct {
	userDB model.UserDB
}

func NewUserDao(userDB model.UserDB) UserDao {
	return &userDaoImp{
		userDB: userDB,
	}
}

func (u *userDaoImp) Create(name, phone, addr string) error {
	err := u.userDB.Create(name, phone, addr)
	if err != nil {
		return err
	}
	return nil
}

//todo 删查改
func (u *userDaoImp) Get(userId int) (*model.User, error) {
	panic("implement me")
}

func (u *userDaoImp) Update(userId int, name, phone, addr string) error {
	panic("implement me")
}

func (u *userDaoImp) List() ([]*model.User, error) {
	res, err := u.userDB.List()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userDaoImp) Query(key string) ([]*model.User, error) {
	panic("implement me")
}
