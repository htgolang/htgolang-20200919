package db

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/define"
	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/utils"
)

// SaveToCSV will flush the define.UserList's contents into a CSV file when
// define.UserList is modified
func SaveToCSV(dbName string, ul *[]define.User) error {
	//dbName = dbName + "." + genFileNameSuffix()
	var subDir = "csv"
	csvFile, err := os.Create(filepath.Join(filepath.Join(dbDir, subDir), dbName))
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
	//dbDir += "/gob"
	var subDir = "gob"
	gob.Register(define.User{})
	gobFile, err := os.Create(filepath.Join(filepath.Join(dbDir, subDir), dbName))
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
func SaveToJSON(dbName string, ul *[]define.User) error {
	//dbDir += "/json"
	var subDir = "json"
	jsonBuffer := new(bytes.Buffer)
	encoder := json.NewEncoder(jsonBuffer)
	encoder.SetIndent("", "\t")
	err := encoder.Encode(ul)
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(filepath.Join(filepath.Join(dbDir, subDir), dbName))
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	_, errjsonW := jsonFile.Write(jsonBuffer.Bytes())
	if errjsonW != nil {
		return errjsonW
	}

	return nil
}

// BackupDB copy a copy of userDB.json/userDB.csv/userDB.gob
func BackupDB(dbDir, subDir, dbName string) error {
	// copy a backup

	fmt.Println("about to copy: ", filepath.Join(filepath.Join(dbDir, subDir), dbName))
	jsonFile, err := os.Open(filepath.Join(filepath.Join(dbDir, subDir), dbName))
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	srcReader := bufio.NewReader(jsonFile)
	backupDBName := dbName + "." + genFileNameSuffix()
	backupFile, err := os.Create(filepath.Join(filepath.Join(dbDir, subDir), backupDBName))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer backupFile.Close()

	_, errW := srcReader.WriteTo(backupFile)
	if errW != nil {
		log.Fatal(errW)
		return errW
	}
	return nil
}

// SaveUsers Wrap the SaveToXXX funcs
func SaveUsers() {
	if SaveFlag == "csv" {
		SaveToCSV(dbNameCSV, &define.UserList)
		BackupDB(dbDir, "csv", dbNameCSV)
		fmt.Println("You already saved via csv, change will be saved automatically.")
		return
	} else if SaveFlag == "gob" {
		SaveToGob(dbNameGob, &define.UserList)
		BackupDB(dbDir, "gob", dbNameGob)
		fmt.Println("You already saved via gob, change will be saved automatically.")
		return
	} else if SaveFlag == "json" {
		SaveToJSON(dbNameJSON, &define.UserList)
		BackupDB(dbDir, "json", dbNameJSON)
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
