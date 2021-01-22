package models

import "time"

type User struct {
	ID        int64
	Name      string
	Sex       bool
	Addr      string
	Tel       string
	Brithday  string
	Passwd    string
	Create_at *time.Time
}

func NewUser(name string, sex bool, addr string, tel string, brithday string, passwd string) (*User, error) {
	createTime := time.Now()
	_, err := time.Parse("2006/01/02", brithday)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:        -1,
		Name:      name,
		Sex:       sex,
		Addr:      addr,
		Tel:       tel,
		Brithday:  brithday,
		Passwd:    passwd,
		Create_at: &createTime,
	}, nil
}
