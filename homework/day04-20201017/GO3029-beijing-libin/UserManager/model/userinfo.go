package model

import (
	"crypto/md5"
)

type UserInfo struct {
	Id      int      `json: "id"`
	Name    string   `json: "string"`
	Pwd     string   `json: "pwd"`
	Enable  bool     `json: "Enable"` //此字段用来判断用户是否可用，即是否被block
	HashPwd [16]byte `json: "hashPwd"`
}

func (userinfo *UserInfo) EditUser(name string, password string) {
	userinfo.Name = name
	userinfo.Pwd = password
	MakeMd5(password, userinfo)
}

//生成userinfo的md5码
func MakeMd5(teststring string, userinfo *UserInfo) {
	var md5Values [16]byte
	md5Values = md5.Sum([]byte(teststring))
	userinfo.HashPwd = md5Values
}

//返回md5后的结果，用于校验
func CheckMd5(s string) [16]byte {
	return md5.Sum([]byte(s))
}

func NewUser(name string, pwd string) *UserInfo {

	return &UserInfo{
		Name: name,
		Pwd:  pwd,
	}
}
