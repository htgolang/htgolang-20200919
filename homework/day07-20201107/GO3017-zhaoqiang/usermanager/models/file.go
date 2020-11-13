package models

import (
	"fmt"
	"os"
)

const confpath string = "users.json"

var File Files = JsonPersistence{confpath}

func init() {
	_, err := os.Stat(confpath)
	if err != nil {
		if os.IsNotExist(err) {
			err := File.InitFile()
			if err != nil {
				fmt.Println("[persistenceFileInit]", err)
				os.Exit(-1)
			}
			fmt.Println("[+] init sucessful")
			return
		}
		fmt.Println("[persistenceFileInit]", err)
		os.Exit(-1)
	}
}
