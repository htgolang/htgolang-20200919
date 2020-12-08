package models

import (
	"encoding/gob"
	"fmt"
	"os"
)

func DbToGob() {
	file, err := os.Create("res/test.gob")
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range UsersDb {
		UserDb = append(UserDb, v)
		// fmt.Println(UserDb)
	}
	enc := gob.NewEncoder(file)
	if err := enc.Encode(UserDb); err != nil {
		fmt.Println(err)
	}

}

func GobToDb() {
	File, _ := os.Open("res/test.gob")
	D := gob.NewDecoder(File)
	D.Decode(&UserDb)
	for _, v := range UserDb {
		UsersDb[v.Name] = v
	}
	// fmt.Println(UsersDb)
}
