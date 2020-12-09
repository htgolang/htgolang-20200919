package model

import "fmt"

type UserSlice struct {
	Length int `json: "length"`
	User   []*UserInfo
}

// 废弃的AddUser 使用注册逻辑来替代Add
// func (ul *UserSlice) AddUser(user *UserInfo) bool {
// 	if user.Enable {
// 		// ul.User[id] = user
// 		ul.User = append(ul.User, user)
// 		return true
// 	}
// 	return false
// }

func (ul *UserSlice) ReturnUser(name string) *UserInfo {
	if ok := ul.FindByNameExist(name); ok {
		return ul.FindUserByName(name)
	}
	return nil
}

func (ul *UserSlice) RemoveUser(name string) {
	for _, v := range ul.User {
		if v.Name == name {
			v.Enable = false
			fmt.Printf("successfully block user %s, his/her attribute is %v", v.Name, v.Enable)
		}
	}
}

func (ul *UserSlice) Showlength() int {
	return len(ul.User)
}
func (ul *UserSlice) FindByNameExist(name string) bool {
	for _, v := range ul.User {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (ul *UserSlice) FindUserByName(name string) (userinfo *UserInfo) {
	for _, v := range ul.User {
		if v.Name == name {
			return v
		}
	}
	return nil

}

func (ul *UserSlice) ShowActiveUser() []*UserInfo {
	fmt.Println("standing by..")
	var activelist []*UserInfo
	for _, v := range ul.User {
		//表示用户可用
		if v.Enable {
			activelist = append(activelist, v)
		}
		// fmt.Println(v)
	}
	return activelist

}

func NewUserSlice() *UserSlice {
	user := NewUser("admin", "123456")
	userslice := &UserSlice{}
	userslice.Length = 1
	userslice.User = append(userslice.User, user)

	return userslice
}
