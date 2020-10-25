package service

import "userManagementV2/model"

type UserService interface {
	//GetUser(userId int) *model.User
	CreateUser(user model.User) error
	ListUser() []*model.User
	//ModifyUser(userId int , user model.User) bool
	//DeleteUser(userId int) bool
	//QueryUser(key string) []*model.User
	//loginVerify(passwd string) bool
}
