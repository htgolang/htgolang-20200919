package models

import (
	"crypto/md5"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

//GenerateElement ...
// generate an element of Users slice, whose type is User
func GenerateElement(ID int, Name string, Tel string, Address string, Birthday string, passwordPlain string) (element *User) {
	element = new(User)
	element.ID, element.Name, element.Tel, element.Address, element.Birthday, element.Password = ID, Name, Tel, Address, func() (ret time.Time) { ret, _ = time.Parse("2006-01-02", Birthday); return }(), md5.Sum([]byte(passwordPlain))
	return
}

//AddElement ...
// Add an element to Users slice, whose type is User
func AddElement(Users *[]User, element User) {
	*Users = append(*Users, element)
}

//RemoveElement ...
// Remove an element from Users slice, whose ID equals the given value
func RemoveElement(Users *[]User, ID int) {
	index := IndexOfElement(Users, ID)
	*Users = append((*Users)[:index], (*Users)[index+1:]...)
}

//ModifyElement ...
// Modify values of keys of specified element from Users slice, whose ID equals the given value
func ModifyElement(Users *[]User, ID int, Name string, Tel string, Address string, Birthday string, passwordPlain string) {
	index := IndexOfElement(Users, ID)
	(*Users)[index].Name = Name
	(*Users)[index].Tel = Tel
	(*Users)[index].Address = Address
	(*Users)[index].Birthday = func() (ret time.Time) { ret, _ = time.Parse("2006-01-02", Birthday); return }()
	if passwordPlain != "" {
		(*Users)[index].Password = md5.Sum([]byte(passwordPlain))
	}
}

//QueryElement ...
// Traverse all User from Users slice, if any of its field contains the given string,
// then append this User to a brand new slice,
// finally return this new slice
func QueryElement(Users *[]User, sub string) (ret []User) {
	ret = make([]User, 0)
	for _, value := range *Users {
		v := reflect.ValueOf(value)
		for i := 0; i < v.NumField(); i++ {
			str := fmt.Sprintf("%v", v.Field(i).Interface())
			if strings.Contains(strings.ToLower(str), sub) {
				ret = append(ret, value)
				break
			}
		}
	}
	return
}

//QueryElementName ...
// Traverse all User from Users slice, if any of its name property equals the given string,
// then append this User to a brand new slice,
// finally return this new slice
func QueryElementName(Users *[]User, sub string) (ret []User) {
	ret = make([]User, 0)
	for _, value := range *Users {
		if value.Name == strings.TrimSpace(sub) {
			ret = append(ret, value)
			break
		}
	}
	return
}

//CheckUserName ...
// Traverse all User from Users slice, if any of its name property equals the given string,
// then return true
// else return false
func CheckUserName(Users *[]User, sub string) (ret bool) {
	ret = false
	for _, value := range *Users {
		if value.Name == strings.TrimSpace(sub) {
			ret = true
			break
		}
	}
	return
}

//CheckUserPassword ...
// Check if input equals the password of specified User.
// Return false if check failed,
// return true if check success.
func CheckUserPassword(user User, password string) (ret bool) {
	if user.Password == md5.Sum([]byte(password)) {
		ret = true
	} else {
		ret = false
	}
	return
}

//IndexOfElement ...
// Returns index of element from Users slice, whose ID equals the given value
// Returns -1 if no ID matches the given value
func IndexOfElement(Users *[]User, ID int) (ret int) {
	ret = -1
	for idx, value := range *Users {
		if ID == value.ID {
			ret = idx
		}
	}
	return
}

//GetMaxID ...
// Returns max ID amoung all elements from Users slice
func GetMaxID(Users *[]User) (ret int) {
	id := -1
	for _, value := range *Users {
		if i := value.ID; i > id {
			id = i
		}
	}
	ret = id
	return
}

//GetMaxIDPlusOne ...
// Returns max ID + 1 amoung all elements from Users slice
func GetMaxIDPlusOne(Users *[]User) (ret int) {
	ret = GetMaxID(Users) + 1
	return
}

//ConvertElementToSlice ...
// Convert the element of Users slices from map to slice
func ConvertElementToSlice(element User) (ret []string) {
	ret = make([]string, 0)
	ret = append(ret, strconv.Itoa(element.ID))
	ret = append(ret, element.Name)
	ret = append(ret, element.Tel)
	ret = append(ret, element.Address)
	ret = append(ret, element.Birthday.Format("2006-01-02"))
	return
}

//PrintElement ...
// Print an element of Users slices in an elegant way
func PrintElement(element User) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Tel", "Address", "Birthday"})
	table.Append(ConvertElementToSlice(element))
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}
