package manager

import (
	"userManagementV2/dao"
	"userManagementV2/model"
)

type UserManager interface {
	CreateUser(user model.User) error
	GetUserById(userId string) (*model.User, error)
	GetUserList() ([]*model.User, error)
	DeleteUserById(userId string) error
	UpdateUser(user model.User) error
	QueryUser(key string) []*model.User
}

type userManagerImp struct {
	userDao dao.UserDao
}

func NewUserManagerImp(userDao dao.UserDao) UserManager {
	return &userManagerImp{
		userDao: userDao,
	}
}

func (u *userManagerImp) CreateUser(user model.User) error {
	err := u.userDao.Create(user.Name, user.Phone, user.Address)
	if err != nil {
		return err
	}
	return nil
}

func (u userManagerImp) GetUserById(userId string) (*model.User, error) {
	panic("implement me")
}

func (u userManagerImp) GetUserList() ([]*model.User, error) {
	res, err := u.userDao.List()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u userManagerImp) DeleteUserById(userId string) error {
	panic("implement me")
}

func (u userManagerImp) UpdateUser(user model.User) error {
	panic("implement me")
}

func (u userManagerImp) QueryUser(key string) []*model.User {
	panic("implement me")
}
