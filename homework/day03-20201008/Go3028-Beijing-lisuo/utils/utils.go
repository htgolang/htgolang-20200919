package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"text/tabwriter"
	"time"

	define "github.com/htgolang/htgolang-20200919/tree/master/homework/day03-20201008/Go3028-Beijing-lisuo/define"
)

// gen a id by UnixNano() who's type is int64
func GenId() (res int64) {
	result := time.Now().UnixNano()
	return result
}

// to verify if a string contains only digits
func JustDigits(s string) bool {
	var a bool = true
	for _, c := range s {
		if c < '0' || c > '9' {
			a = false
			break
		}
	}
	return a
}

// read content from standard input
func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	return line
}

// show a user based on Id
func ShowUser(Id int64) {
	for _, userMap := range define.UserList {
		if val, ok := userMap[Id]; ok {
			//fmt.Println(i, Id, val)
			//fmt.Println(val.Name, val.Address, val.Phone)
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0|tabwriter.Debug)
			s := strconv.FormatInt(Id, 10)
			fmt.Fprintln(w, "|"+s+"\t"+val.Name+"\t"+val.Phone+"\t"+val.Address+" |")
			w.Flush()
		}
	}
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
func NameFindUser(Name string) map[int64]define.User {
	var user map[int64]define.User
	var id int64
	for _, userMap := range define.UserList {
		for k, v := range userMap {
			if v.Name == Name {
				user = map[int64]define.User{k: define.User{v.Name, v.Address, v.Phone}}
				id = k
			}
		}
	}
	if string(id) == "" {
		fmt.Println("No such user.")
	}
	return user
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

// clear the console
var clear map[string]func() //create a map for storing clear funcs
func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
