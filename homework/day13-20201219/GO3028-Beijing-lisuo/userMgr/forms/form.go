package forms

import (
	"time"
)

type UserForm struct {
	ID      int64     `form:"id"`
	Name    string    `form:"name"`
	Sex     int       `form:"sex"`
	Address string    `form:"address"`
	Cell    string    `form:"cell"`
	Born    time.Time `form:"born"`
	Passwd  string    `form:"passwd"`
}
