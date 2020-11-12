package models

import (
	"crypto/md5"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

//User ...
//Global data structure of User element.
type User struct {
	UserID   int       `json:"id"`
	Name     string    `json:"name" boltholdIndex:"Name"`
	Tel      string    `json:"tel"`
	Address  string    `json:"address"`
	Birthday time.Time `json:"birthday"`
	Password [16]byte  `json:"password"`
}

//Users ...
//Global Data Structure, for each element of Users slice,
//in other words, User, has the same data structure below:
/*
	ID int
	Name string
	Tel string
	Address string
	Birthday time.Time
	password [16]byte
*/
var Users []User

//CheckUserPassword ...
// Check if input equals the password of specified User.
// Return false if check failed,
// return true if check success.
func (u User) CheckUserPassword(password string) (ret bool) {
	if u.Password == md5.Sum([]byte(password)) {
		ret = true
	} else {
		ret = false
	}
	return
}

//ConvertElementToSlice ...
// Convert the element of Users slices from map to slice
func (u User) ConvertElementToSlice() (ret []string) {
	ret = make([]string, 0)
	ret = append(ret, strconv.Itoa(u.UserID))
	ret = append(ret, u.Name)
	ret = append(ret, u.Tel)
	ret = append(ret, u.Address)
	ret = append(ret, u.Birthday.Format("2006-01-02"))
	return
}

//PrintElement ...
// Print an element of Users slices in an elegant way
func (u User) PrintElement() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Tel", "Address", "Birthday"})
	table.Append(u.ConvertElementToSlice())
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}
