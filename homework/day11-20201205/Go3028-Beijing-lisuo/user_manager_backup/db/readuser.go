package db

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/define"
)

// Read read Users from dbAbsFile, append them to a UserList
func Read(dbName string, ul *[]define.User) {
	switch dbName {
	case dbNameCSV:
		readFromCSV(dbNameCSV, ul)
	case dbNameGob:
		readFromGob(dbNameGob, ul)
	case dbNameJSON:
		readFromJSON(dbNameJSON, ul)
	default:
		readFromCSV(dbNameCSV, ul)
	}
}

func readFromCSV(dbName string, ul *[]define.User) {
	var subDir = "csv"
	csvFile, err := os.Open(filepath.Join(filepath.Join(dbDir, subDir), dbName))
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		userID, _ := strconv.ParseInt(line[0], 10, 64)
		tmpUsers = append(tmpUsers,
			define.User{ID: userID,
				Name:    line[1],
				Cell:    line[2],
				Address: line[3],
				Born: func() time.Time {
					t, _ := time.Parse("2006.01.02", line[4])
					return t
				}(),
				Passwd: line[5],
			})
	}
	// assign the read users to define.UserList
	*ul = tmpUsers
	// purge the tmpUsers
	tmpUsers = []define.User{}
}

func readFromGob(dbName string, ul *[]define.User) {
	var subDir = "gob"
	gobFile, err := os.Open(filepath.Join(filepath.Join(dbDir, subDir), dbName))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer gobFile.Close()

	gobDecoder := gob.NewDecoder(gobFile)
	if err := gobDecoder.Decode(ul); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//fmt.Print(ul)
}

func readFromJSON(dbName string, ul *[]define.User) error {
	var subDir = "json"
	jsonFile, err := os.Open(filepath.Join(filepath.Join(dbDir, subDir), dbName))
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonReader := bufio.NewReader(jsonFile)
	jsonBuffer := new(bytes.Buffer)
	jsonReader.WriteTo(jsonBuffer)
	decoder := json.NewDecoder(jsonBuffer)

	if err := decoder.Decode(ul); err != nil {
		return err
	}
	return nil

}

// ReadUsers wrap Read
func ReadUsers() {
	Read(SaveFlag, &define.UserList)
}
