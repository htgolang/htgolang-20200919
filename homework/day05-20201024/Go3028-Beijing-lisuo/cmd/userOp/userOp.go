package userOp

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/define"
	"github.com/olekukonko/tablewriter"
)

// show a user based on Id
func ShowUser(Id int64) {
	t := tablewriter.NewWriter(os.Stdout)
	//t.SetAutoFormatHeaders(false)
	//t.SetAutoWrapText(false)
	//t.SetReflowDuringAutoWrap(false)
	//t.SetHeader([]string{"ID", "Name", "Phone", "Location"})
	for _, userMap := range define.UserList {
		if v, ok := userMap[Id]; ok {
			s := strconv.FormatInt(Id, 10)
			t.Append([]string{s, v.Name, v.Phone, v.Address})
		}
	}
	t.Render()
}

// find user based on Id
func IdFindUser(UserList []map[int64]define.User, Id int64) define.User {
	var user define.User
	for _, userMap := range UserList {
		if _, ok := userMap[Id]; ok {
			user = userMap[Id]
		}
	}
	return user
}

// find user based on Name
//func NameFindUser(Name string) map[int64]define.User {
//	var user map[int64]define.User
//	var id int64
//	for _, userMap := range define.UserList {
//		for k, v := range userMap {
//			if v.Name == Name {
//				user = map[int64]define.User{k: define.User{v.Name, v.Address, v.Phone}}
//				id = k
//			}
//		}
//	}
//	if string(id) == "" {
//		fmt.Println("No such user.")
//	}
//	return user
//}

// find user based on Name
func NameFindUser(ul *define.UserList, Name string) (define.User, error) {
	for i, user := range (*ul) {
			if user.name == Name {
				return user, nil
			}
		}
			err := errors.New("There's no such user.")
			return define.User{}, err
	}
}


// del user based on Id
func IdDelUser(user *[]map[int64]define.User, id int64) {
	for i, u := range *user {
		for k, _ := range u {
			if int64(k) == id {
				*user = append(define.UserList[:i], define.UserList[i+1:]...)
			}
		}
	}
}

// del user based on Name
func NameDelUser(user *[]map[int64]define.User, name string) {
	for i, u := range *user {
		for _, v := range u {
			if v.Name == name {
				if i == len(*user) {
					*user = append(define.UserList[:i], define.UserList[i:]...)
					return
				}
				*user = append(define.UserList[:i], define.UserList[i+1:]...)
			}
		}
	}
}

// modify user based on Id
func IdModUser(user *[]map[int64]define.User, id int64) define.User {
	var iname, iaddress, iphone string
	var newUser define.User
	for _, u := range *user {
		for k, v := range u {
			if int64(k) == id {
				fmt.Printf("Input new Name [%v]: ", v.Name)
				iname = Read()
				newUser.Name = iname
				if iname == "" {
					newUser.Name = v.Name
				}
				fmt.Printf("Input new Address [%v]: ", v.Address)
				iaddress = Read()
				newUser.Address = iaddress
				if iaddress == "" {
					newUser.Address = v.Address
				}
				fmt.Printf("Input new Phone [%v]: ", v.Phone)
				iphone = Read()
				// make sure the phone number contains only pure digits
				for JustDigits(iphone) == false {
					fmt.Print("Please input a legal phone number: \n> ")
					iphone = Read()
					if JustDigits(iphone) == true {
						break
					}
				}
				newUser.Phone = iphone
				if iphone == "" {
					newUser.Phone = v.Phone
				}
				u[k] = newUser
				fmt.Printf("Modified user is: %v:%v\n", k, newUser)
			}
		}
	}
	return newUser
}

// modify user based on Name
func NameModUser(user *[]map[int64]define.User, name string) {
	var iname, iaddress, iphone string
	var newUser define.User
	for _, u := range *user {
		for k, v := range u {
			if v.Name == name {
				fmt.Printf("Input new Name [%v]: ", v.Name)
				iname = Read()
				newUser.Name = iname
				if iname == "" {
					newUser.Name = v.Name
				}
				fmt.Printf("Input new Address [%v]: ", v.Address)
				iaddress = Read()
				newUser.Address = iaddress
				if iaddress == "" {
					newUser.Address = v.Address
				}
				fmt.Printf("Input new Phone [%v]: ", v.Phone)
				iphone = Read()
				// make sure the phone number contains only pure digits
				for JustDigits(iphone) == false {
					fmt.Print("Please input a legal phone number: \n> ")
					iphone = Read()
					if JustDigits(iphone) == true {
						break
					}
				}
				newUser.Phone = iphone
				if iphone == "" {
					newUser.Phone = v.Phone
				}
				u[k] = newUser
				fmt.Printf("Modified user is: %v:%v\n", k, newUser)
			}
		}
	}
}