package userop

import (
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/utils"
	"github.com/olekukonko/tablewriter"
)

// NewUser make a new user contains user's info
func NewUser(id int64, name, cell, address, born, passwd string) define.User {
	return define.User{
		Id:      id,
		Name:    name,
		Cell:    cell,
		Address: address,
		Born: func() time.Time {
			t, _ := time.Parse("2006.01.02", born)
			return t
		}(),
		Passwd: md5.Sum([]byte(passwd)),
	}
}

// Init add some users to define.UserList
func Init(ul *[]define.User) {
	user0 := NewUser(0, "admin", "18811992299", "HaidianDistrict,BeijingXinParkRestaurants,BeixiaguanSubdistrict,HaidianDistrict,China", time.Now().Format("2006.01.02"), "qwert")
	user1 := NewUser(1, "jack ma", "18800009999", "Hangzhou, China", time.Now().Format("2006.01.02"), "hello")
	user3 := NewUser(3, "steve", "18800002222", "Mars", time.Now().Format("2006.01.02"), "hi")
	define.UserList = append(*ul, user0, user1, user3)
	fmt.Printf("user %v added\n", user0.Name)
	fmt.Printf("user %v added\n", user1.Name)
	fmt.Printf("user %v added\n", user3.Name)
}

// ShowUser show a user based on Id
func ShowUser(id int64) {
	t := tablewriter.NewWriter(os.Stdout)
	for _, user := range define.UserList {
		if user.Id == id {
			s := strconv.FormatInt(id, 10)
			p := fmt.Sprintf("%x", user.Passwd)
			t.Append([]string{s, user.Name, user.Cell, user.Address, user.Born.Format("2006.01.02"), p})
			//md5.Sum([]byte(user.Passwd))})
		}
	}
	t.Render()
}

// find user based on Id
func IdFindUser(ul *[]define.User, id int64) (define.User, error) {
	for _, user := range *ul {
		if user.Id == id {
			return user, nil
		}
	}
	err := errors.New("no such user")
	return define.User{}, err
}

// find user based on Name
func NameFindUser(ul *[]define.User, Name string) (define.User, error) {
	for _, user := range *ul {
		if user.Name == Name {
			return user, nil
		}
	}
	err := errors.New("no such user")
	return define.User{}, err
}

// del user based on Id
func IdDelUser(ul *[]define.User, id int64) {
	for i, user := range *ul {
		if int64(user.Id) == id {
			*ul = append(define.UserList[:i], define.UserList[i+1:]...)
		}
	}
}

// del user based on Name
func NameDelUser(ul *[]define.User, name string) {
	for i, user := range *ul {
		if user.Name == name {
			if i == len(*ul) {
				*ul = append(define.UserList[:i], define.UserList[i:]...)
				return
			}
			*ul = append(define.UserList[:i], define.UserList[i+1:]...)
		}
	}
}

// get max id
func GetMaxID(ul *[]define.User) int64 {
	var MaxID int64 = -1
	for _, user := range *ul {
		var i int64 = user.Id
		if i > MaxID {
			MaxID = i
		}
	}
	return MaxID
}

// modify user based on Id
func IdModUser(ul *[]define.User, id int64) define.User {
	var iname, iaddress, cell string
	var ipasswd string
	var newUser define.User
	for i, user := range *ul {
		if user.Id == id {
			fmt.Printf("Input new Name [%v]: ", user.Name)
			iname = utils.Read()
			newUser.Name = iname
			if iname == "" {
				newUser.Name = user.Name
			}
			fmt.Printf("Input new Address [%v]: ", user.Address)
			iaddress = utils.Read()
			newUser.Address = iaddress
			if iaddress == "" {
				newUser.Address = user.Address
			}
			fmt.Printf("Input new Phone [%v]: ", user.Cell)
			cell = utils.Read()
			// make sure the phone number contains only pure digits
			for utils.JustDigits(cell) == false {
				fmt.Print("Please input a legal phone number: \n> ")
				cell = utils.Read()
				if utils.JustDigits(cell) == true {
					break
				}
			}
			newUser.Cell = cell
			if cell == "" {
				newUser.Cell = user.Cell
			}
			fmt.Printf("Input new passwd [%v]: ", user.Passwd)
			newUser.Passwd = utils.GenPasswd()
			if ipasswd == "" {
				newUser.Passwd = user.Passwd
			}

			(*ul)[i] = newUser
			fmt.Printf("Modified user is: %v\n", (*ul)[i])
		}
	}
	return newUser
}

// to recap =================================
// modify user based on Name
//func NameModUser(user *[]map[int64]define.User, name string) {
//	var iname, iaddress, iphone string
//	var newUser define.User
//	for _, u := range *user {
//		for k, v := range u {
//			if v.Name == name {
//				fmt.Printf("Input new Name [%v]: ", v.Name)
//				iname = utils.Read()
//				newUser.Name = iname
//				if iname == "" {
//					newUser.Name = v.Name
//				}
//				fmt.Printf("Input new Address [%v]: ", v.Address)
//				iaddress = utils.Read()
//				newUser.Address = iaddress
//				if iaddress == "" {
//					newUser.Address = v.Address
//				}
//				fmt.Printf("Input new Phone [%v]: ", v.Phone)
//				iphone = utils.Read()
//				// make sure the phone number contains only pure digits
//				for utils.JustDigits(iphone) == false {
//					fmt.Print("Please input a legal phone number: \n> ")
//					iphone = utils.Read()
//					if utils.JustDigits(iphone) == true {
//						break
//					}
//				}
//				newUser.Phone = iphone
//				if iphone == "" {
//					newUser.Phone = v.Phone
//				}
//				u[k] = newUser
//				fmt.Printf("Modified user is: %v:%v\n", k, newUser)
//			}
//		}
//	}
//}
//
