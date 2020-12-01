package utils

import (
	"net/http"
	"sort"
	"usermanage/model"
)

// 获取当前最大用户Id
func GetMaxId() int {
	if len(UsersList) > 0 {
		sort.Sort(UsersList)
		return UsersList[len(UsersList) - 1].Id
	}
	return 0
}

// 判断一个用户是否存在,存在时返回切片下标
func IsUserExists(name string) (int, bool) {
	for i, user := range UsersList {
		if name == user.Name {
			return i, true
		}
	}
	return 0, false
}

// 判断两次密码是否一致
func IsPasswdSame(r *http.Request) (bool) {
	return r.FormValue("password") == r.FormValue("confirm")
}

// 根据用户ID获取用户信息
func GetUserInfoById(id int) *model.User {
	for _, v := range UsersList {
		if v.Id == id {
			return v
		}
	}
	return &model.User{}
}