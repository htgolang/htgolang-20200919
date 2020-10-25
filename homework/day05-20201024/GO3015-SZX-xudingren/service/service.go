package service

import "GO3015-SZX-xudingren/model"

type Service interface {
	GetUser(userId int) *model.User
	AddUser(user model.User) bool
	ModifyUser(userId int , user model.User) bool
	DeleteUser(userId int) bool
	ListUser() []*model.User
	QueryUser(key string) []*model.User
	loginVerify(passwd string) bool
}