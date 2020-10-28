package models

import (
	"errors"
)

//FindUserByID 查找用户 并返回用户信息
func FindUserByID(id int) (User, error) {
	for _, user := range users {
		if user.id == id {
			return user, nil
		}
	}
	return users[0], errors.New("user not exist")
}

//DelUser 删除用户
func DelUser(u User) {
	tmpUsers := make([]User, 0, len(users))
	for _, user := range users {
		if user.id != u.id {
			tmpUsers = append(tmpUsers, user)
		}
	}

	users = tmpUsers
}
