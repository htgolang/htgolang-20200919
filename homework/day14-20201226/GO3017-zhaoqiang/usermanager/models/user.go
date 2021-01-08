package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64     `form:"id" orm:"column(id);pk;auto"`
	Name     string    `form:"name" orm:"size(25)"`
	PassWord string    `form:"password" orm:"column(password);size(70)"`
	Sex      bool      `form:"sex" orm:"null"`
	Addr     string    `form:"addr" orm:"type(text);null"`
	Tel      string    `form:"tel" orm:"size(20);null"`
	Admin    bool      `orm:"column(admin);default(0)"`
	Brithday time.Time `form:"brithday" orm:"type(date)"`
	CreateAt time.Time `orm:"auto_now_add"`
	UpdateAt time.Time `orm:"auto_now"`
}

func (u *User) TableName() string {
	return "userorm"
}

func (u *User) TableIndex() [][]string {
	return [][]string{
		{"Brithday"},
		{"Name", "PassWord", "Addr"},
	}
}

func NewUser(name string, pass string, sex bool, addr string, tel string, brithday string) (*User, error) {
	brith, err := time.Parse("2006-01-02", brithday)
	if err != nil {
		return nil, err
	}
	cryptopass, err := PassCrypto(pass)
	if err != nil {
		return nil, errors.New("passwd crypto false")
	}
	return &User{
		Name:     name,
		PassWord: cryptopass,
		Sex:      sex,
		Addr:     addr,
		Tel:      tel,
		Brithday: brith,
	}, nil
}

func PassCrypto(pass string) (string, error) {
	hasher, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
	if err != nil {
		return "", err
	}
	return string(hasher), nil
}
