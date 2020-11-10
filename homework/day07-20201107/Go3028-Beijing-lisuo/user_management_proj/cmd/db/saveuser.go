package db

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day07-20201107/Go3028-Beijing-lisuo/user_management_proj/define"
)

// Save will flush the define.UserList's contents into a file periodically
func Save(ul *[]define.User) {
	csvFile, err := os.Create(filepath.Join(dbDir, dbName))
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(csvFile)
	for _, user := range *ul {
		writer.Write([]string{strconv.FormatInt(user.ID, 10), user.Name, user.Cell, user.Address,
			user.Born.Format("2006.01.02"), user.Passwd})
	}
	writer.Flush()
}

// SaveUsers Wrap the Save func
func SaveUsers() {
	Save(&define.UserList)
}
