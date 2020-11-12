package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

//MemoryUsersEntity ...
type MemoryUsersEntity struct {
	Users             *[]User
	PersistentStorage Storage
}

//MUE ...
var MUE *MemoryUsersEntity

//PersistentStorage ...
var PersistentStorage Storage

//GenerateElement ...
// generate an element of Users slice, whose type is User
func GenerateElement(UserID int, Name string, Tel string, Address string, Birthday string, passwordPlain string) (element *User) {
	element = new(User)
	element.UserID, element.Name, element.Tel, element.Address, element.Birthday, element.Password = UserID, Name, Tel, Address, func() (ret time.Time) { ret, _ = time.Parse("2006-01-02", Birthday); return }(), md5.Sum([]byte(passwordPlain))
	return
}

//GetElement ...
// get an element of Users slice, whose ID equals the given value
func (mue *MemoryUsersEntity) GetElement(UserID int) (ret User, err error) {
	err = errors.New("User not found error")
	for _, value := range *mue.Users {
		if UserID == value.UserID {
			ret = value
			err = nil
			break
		}
	}
	return
}

//AddElement ...
// Add an element to Users slice, whose type is User
func (mue *MemoryUsersEntity) AddElement(element User) {
	*mue.Users = append(*mue.Users, element)
}

//RemoveElement ...
// Remove an element from Users slice, whose ID equals the given value
func (mue *MemoryUsersEntity) RemoveElement(UserID int) {
	index := mue.IndexOfElement(UserID)
	*mue.Users = append((*mue.Users)[:index], (*mue.Users)[index+1:]...)
}

//ModifyElement ...
// Modify values of keys of specified element from Users slice, whose ID equals the given value
func (mue *MemoryUsersEntity) ModifyElement(UserID int, Name string, Tel string, Address string, Birthday string, passwordPlain string) {
	index := mue.IndexOfElement(UserID)
	(*mue.Users)[index].Name = Name
	(*mue.Users)[index].Tel = Tel
	(*mue.Users)[index].Address = Address
	(*mue.Users)[index].Birthday = func() (ret time.Time) { ret, _ = time.Parse("2006-01-02", Birthday); return }()
	if passwordPlain != "" {
		(*mue.Users)[index].Password = md5.Sum([]byte(passwordPlain))
	}
}

//QueryElement ...
// Traverse all User from Users slice, if any of its field contains the given string,
// then append this User to a brand new slice,
// finally return this new slice
func (mue *MemoryUsersEntity) QueryElement(sub string) (ret []User) {
	ret = make([]User, 0)
	for _, value := range *mue.Users {
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
func (mue *MemoryUsersEntity) QueryElementName(sub string) (ret []User) {
	ret = make([]User, 0)
	for _, value := range *mue.Users {
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
func (mue *MemoryUsersEntity) CheckUserName(sub string) (ret bool) {
	ret = false
	for _, value := range *mue.Users {
		if value.Name == strings.TrimSpace(sub) {
			ret = true
			break
		}
	}
	return
}

//IndexOfElement ...
// Returns index of element from Users slice, whose ID equals the given value
// Returns -1 if no ID matches the given value
func (mue *MemoryUsersEntity) IndexOfElement(ID int) (ret int) {
	ret = -1
	for idx, value := range *mue.Users {
		if ID == value.UserID {
			ret = idx
		}
	}
	return
}

//GetMaxID ...
// Returns max ID amoung all elements from Users slice
func (mue *MemoryUsersEntity) GetMaxID() (ret int) {
	id := -1
	for _, value := range *mue.Users {
		if i := value.UserID; i > id {
			id = i
		}
	}
	ret = id
	return
}

//GetMaxIDPlusOne ...
// Returns max ID + 1 amoung all elements from Users slice
func (mue *MemoryUsersEntity) GetMaxIDPlusOne() (ret int) {
	ret = mue.GetMaxID() + 1
	return
}
