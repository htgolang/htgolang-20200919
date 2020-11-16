package db

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day07-20201107/Go3028-Beijing-lisuo/user_management_proj/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day07-20201107/Go3028-Beijing-lisuo/user_management_proj/utils"
)

// SaveToCSV will flush the define.UserList's contents into a CSV file when
// define.UserList is modified
func SaveToCSV(dbName string, ul *[]define.User) error {
	csvFile, err := os.Create(filepath.Join(dbDir, dbName))
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)
	for _, user := range *ul {
		writer.Write([]string{strconv.FormatInt(user.ID, 10), user.Name, user.Cell, user.Address,
			user.Born.Format("2006.01.02"), user.Passwd})
	}
	writer.Flush()
	return writer.Error()
}

// SaveToGob encode the UserList to a gob file
func SaveToGob(dbName string, ul *[]define.User) error {
	gob.Register(define.User{})
	gobFile, err := os.Create(filepath.Join(dbDir, dbName))
	if err != nil {
		//log.Fatal(err)
		return err
	}
	defer gobFile.Close()
	gobEncoder := gob.NewEncoder(gobFile)
	if err := gobEncoder.Encode(&ul); err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

// SaveToJSON encode the UserList to a json file
func SaveToJSON(dbName string, ul *[]define.User) (int, error) {
	jsonBuffer := new(bytes.Buffer)
	encoder := json.NewEncoder(jsonBuffer)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(ul)
	if err != nil {
		return 0, err
	}

	jsonFile, err := os.Create(filepath.Join(dbDir, dbName))
	if err != nil {
		return 0, err
	}

	numberOfBytesWritten, err := jsonFile.Write(jsonBuffer.Bytes())
	if err != nil {
		return 0, err
	}
	return numberOfBytesWritten, nil
}

// SaveUsers Wrap the SaveToXXX funcs
func SaveUsers() {
	if SaveFlag == "csv" {
		SaveToCSV(dbNameCSV, &define.UserList)
		fmt.Println("You already saved via csv, change will be saved automatically.")
		return
	} else if SaveFlag == "gob" {
		SaveToGob(dbNameGob, &define.UserList)
		fmt.Println("You already saved via gob, change will be saved automatically.")
		return
	} else if SaveFlag == "json" {
		SaveToJSON(dbNameJSON, &define.UserList)
		fmt.Println("You already saved via json, change will be saved automatically.")
		return
	} else {
		utils.Message("Please choose which format to save users: ")
		for k, v := range dbNameList {
			fmt.Printf("%v : %v\n", k, v)
		}
		input := utils.Read()
		switch input {
		case "csv":
			SaveToCSV(dbNameCSV, &define.UserList)
			SaveFlag = "csv"
		case "gob":
			SaveToGob(dbNameGob, &define.UserList)
			SaveFlag = "gob"
		case "json":
			SaveToJSON(dbNameJSON, &define.UserList)
			SaveFlag = "json"
		}
	}
}
