package models

import "errors"

//Authentication 验证用户输入的用户名是否重复
func Authentication(u User) bool {
	for _, user := range users {
		if user.name == u.name {
			return false
		}
	}
	return true
}

//Authentication1 asdf
func Authentication1(olduser User, newuser User) (bool, error) {
	for _, user := range users {
		if user == olduser {
			continue
		}
		if user.id == newuser.id {
			return false, errors.New("user id already")
		}
		if user.name == newuser.name {
			return false, errors.New("user name already")
		}
	}
	return true, nil
}
