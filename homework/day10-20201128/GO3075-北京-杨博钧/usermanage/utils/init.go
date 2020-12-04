package utils

import "usermanage/model"

// 用户列表
type UserList []*model.User
var UsersList UserList
//定义方法满足sort接口
func (this UserList) Len() int {
	return len(this)
}

func (this UserList) Less(i, j int) bool {
	return this[i].Id < this[j].Id
}

func (this UserList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
