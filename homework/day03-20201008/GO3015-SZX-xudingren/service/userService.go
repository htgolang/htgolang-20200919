package service

import (
	"GO3015-SZX-xudingren/model"
	"errors"
	"fmt"
	"regexp"
)

//具体业务逻辑处理：增删查改

type UserService struct {
	users   []model.User
	userNum int //用户数
	curId   int //当前用户ID
}

//工厂函数
func NewUserService() *UserService {
	userService := &UserService{}
	userService.userNum = 1
	userService.curId = 1
	user := model.User{
		Id:      1,
		Name:    "zhangsan",
		Phone:   "13788888888",
		Address: "dz",
	}
	userService.users = append(userService.users, user)
	return userService
}

func (s *UserService) GetUserNum() int {
	return s.userNum
}

func (s *UserService) HasUser(userId int) (int, error) {
	for i, v := range s.users {
		if v.Id == userId {
			return i, nil
		}
	}
	return -1, errors.New("查无此用户")
}

func (s *UserService) GetUser(idx int) []model.User {
	return []model.User{s.users[idx]}
}

func (s *UserService) Add(user model.User) bool {
	s.userNum++
	s.curId++
	user.Id = s.curId
	s.users = append(s.users, user)
	return true
}

func (s *UserService) Modify(idx int, user model.User) bool {
	s.users[idx].Name = user.Name
	s.users[idx].Phone = user.Phone
	s.users[idx].Address = user.Address
	return true
}

func (s *UserService) Delete(idx int) bool {
	s.users = append(s.users[:idx], s.users[idx+1:]...)
	s.userNum--
	return true
}

func (s *UserService) List() []model.User {
	return s.users
}

func (s *UserService) Query(k string) []model.User {
	re := regexp.MustCompile(k)
	if re == nil {
		fmt.Println("regexp err")
		return nil
	}
	users := []model.User{}
	for _, v := range s.users {
		if re.MatchString(v.Name) {
			users = append(users, v)
			continue
		}
		if re.MatchString(v.Address) {
			users = append(users, v)
			continue
		}
		if re.MatchString(v.Phone) {
			users = append(users, v)
			continue
		}
	}
	return users
}
