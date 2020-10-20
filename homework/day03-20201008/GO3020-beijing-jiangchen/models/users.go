package models

import (
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

//Users ...
//Global Data Structure, for each element of Users slice,
//in other words, User, has the same data structure below:
/*
	ID
	Name
	Contact
	Address
*/
var Users []map[string]string

func init() {
	Users = make([]map[string]string, 0)
	Users = append(Users, GenerateElement("1001", "Alice", "+1 4406665321", "2426 Wildwood Street, Medina, Ohio"))
	Users = append(Users, GenerateElement("1002", "Norman", "+1 6789143737", "4548 Davis Street, Norcross, Georgia"))
	Users = append(Users, GenerateElement("1003", "Connie", "+1 2184173411", "1485 Laurel Lee, Forest Lake, Minnesota"))
	Users = append(Users, func(ID string, Name string, Contact string, Address string) map[string]string {
		element := make(map[string]string)
		element["ID"], element["Name"], element["Contact"], element["Address"] = ID, Name, Contact, Address
		return element
	}("1004", "David", "+1 8455462309", "3627 Camden Place, Poughkeepsie, New York"))
}

//GenerateElement ...
// generate an element of Users slice, whose type is map[string]string
func GenerateElement(ID string, Name string, Contact string, Address string) (element map[string]string) {
	element = make(map[string]string)
	element["ID"], element["Name"], element["Contact"], element["Address"] = ID, Name, Contact, Address
	return
}

//AddElement ...
// Add an element to Users slice, whose type is map[string]string
func AddElement(Users *[]map[string]string, element map[string]string) {
	*Users = append(*Users, element)
}

//RemoveElement ...
// Remove an element from Users slice, whose ID equals the given value
func RemoveElement(Users *[]map[string]string, ID string) {
	index := IndexOfElement(Users, ID)
	*Users = append((*Users)[:index], (*Users)[index+1:]...)
}

//ModifyElement ...
// Modify values of keys of specified element from Users slice, whose ID equals the given value
func ModifyElement(Users *[]map[string]string, ID string, Name string, Contact string, Address string) {
	index := IndexOfElement(Users, ID)
	(*Users)[index]["Name"] = Name
	(*Users)[index]["Contact"] = Contact
	(*Users)[index]["Address"] = Address
}

//QueryElement ...
// Traverse all User from Users slice, if any of its field contains the given string,
// then append this User to a brand new slice,
// finally return this new slice
func QueryElement(Users *[]map[string]string, sub string) (ret []map[string]string) {
	ret = make([]map[string]string, 0)
	for _, value := range *Users {
		for _, v := range value {
			if strings.Contains(strings.ToLower(v), sub) {
				ret = append(ret, value)
				break
			}
		}
	}
	return
}

//IndexOfElement ...
// Returns index of element from Users slice, whose ID equals the given value
// Returns -1 if no ID matches the given value
func IndexOfElement(Users *[]map[string]string, ID string) (ret int) {
	ret = -1
	for idx, value := range *Users {
		if ID == value["ID"] {
			ret = idx
		}
	}
	return
}

//GetMaxID ...
// Returns max ID amoung all elements from Users slice
func GetMaxID(Users *[]map[string]string) (ret string) {
	id := -1
	for _, value := range *Users {
		if i, _ := strconv.Atoi(value["ID"]); i > id {
			id = i
		}
	}
	ret = strconv.Itoa(id)
	return
}

//GetMaxIDPlusOne ...
// Returns max ID + 1 amoung all elements from Users slice
func GetMaxIDPlusOne(Users *[]map[string]string) (ret string) {
	MaxID := GetMaxID(Users)
	i, _ := strconv.Atoi(MaxID)
	ret = strconv.Itoa(i + 1)
	return
}

//ConvertElementToSlice ...
// Convert the element of Users slices from map to slice
func ConvertElementToSlice(element map[string]string) (ret []string) {
	ret = make([]string, 0)
	ret = append(ret, element["ID"])
	ret = append(ret, element["Name"])
	ret = append(ret, element["Contact"])
	ret = append(ret, element["Address"])
	return
}

//PrintElement ...
// Print an element of Users slices in an elegant way
func PrintElement(element map[string]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Name", "Contact", "Address"})
	table.Append(ConvertElementToSlice(element))
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Render()
}
