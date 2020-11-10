package db

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day07-20201107/Go3028-Beijing-lisuo/user_management_proj/define"
)

// Read read Users from dbAbsFile, append them to a UserList
func Read(dbName string, ul *[]define.User) {

	//path, err := filepath.Abs(base)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//absDir = filepath.Dir(path)
	//dbLocation = filepath.Join(absDir, dbDir)
	//dbAbsFile = filepath.Join(dbLocation, dbName)

	csvFile, err := os.Open(filepath.Join(dbDir, dbName))
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
	(*ul) = tmpUsers
	// purge the tmpUsers
	tmpUsers = []define.User{}
}

// ReadUsers wrap Read
func ReadUsers() {
	Read(dbName, &define.UserList)
}
