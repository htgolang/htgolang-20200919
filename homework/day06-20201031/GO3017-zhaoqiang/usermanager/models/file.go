package models

import (
	"fmt"
	"os"
)

const filepath string = "users.csv"

var File Files = GobPersistence{filepath}

func init() {
	_, err := os.Stat(filepath)
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
