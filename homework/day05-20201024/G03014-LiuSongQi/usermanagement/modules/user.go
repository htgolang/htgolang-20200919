package modules

import (
	"time"
	"usermanagement/utils"
)

type UserInfo struct {
	Id       int
	Name     string
	Address  string
	Tel      string
	Birthday time.Time
	Password string
	// Online   bool
}

type UserManager struct {
	Users []UserInfo
}

func NewUserManager() *UserManager {
	return &UserManager{
		Users: make([]UserInfo, 0),
	}
}

func (u *UserManager) GetId() int {
	id := 0
	for _, v := range u.Users {
		if v.Id > id {
			id = v.Id
		}
	}
	return id + 1
}

func (m *UserManager) SetDefaultUser() {
	var defaltUser = UserInfo{
		Id:       0,
		Name:     "admin",
		Password: "E10ADC3949BA59ABBE56E057F20F883E",
		Address:  "上海市中山公园",
		Tel:      "1212121",
		Birthday: utils.TimeConversion("1992-10-10"),
	}
	m.Users = append(m.Users, defaltUser)
}

func (u *UserManager) AddUser(user *UserInfo) {
	u.Users = append(u.Users, *user)
}

func (u *UserManager) Login(username, passwd string) bool {
	for _, user := range u.Users {
		if user.Name == username && user.Password == passwd {
			return true
		}
	}
	return false
}

func (u *UserManager) FindUserByName(username string) bool {
	for _, user := range u.Users {
		if user.Name == username {
			return true
		}
	}
	return false
}

func (u *UserManager) FindUserById(id int) (int, bool) {
	for idx, user := range u.Users {
		if user.Id == id {
			return idx, true
		}
	}
	return 0, false
}

func (u *UserManager) DeleteUser(idx int) {
	u.Users = append(u.Users[:idx], u.Users[idx+1:]...)
}

func (u *UserManager) ModifyUser(idx int, user *UserInfo) {
	u.Users[idx] = *user
}

func (u *UserManager) QueryUser(name string) *UserInfo {
	for _, user := range u.Users {
		if user.Name == name {
			return &user
		}
	}
	return nil
}
