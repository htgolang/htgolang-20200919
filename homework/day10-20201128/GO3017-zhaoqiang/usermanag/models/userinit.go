package models

import (
	"fmt"
	"os"
)

const confname string = "users.json"

var File Files = NewJSONPersistence(confname)

func init() {
	_, err := os.Stat("../data/" + confname)
	if err != nil {
		if os.IsNotExist(err) {
			err := File.InitFile()
			if err != nil {
				fmt.Println("[persistenceFile--File.InitFile]", err)
				os.Exit(-1)
			}
			fmt.Println("[+]init sucessful")
			return
		}
		fmt.Println("[persistenceFileInit]", err)
		os.Exit(-1)
	}
}
