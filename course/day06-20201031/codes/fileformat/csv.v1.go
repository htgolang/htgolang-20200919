package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type User struct {
	ID   int
	Name string
}

func main() {
	users := []User{
		{1, "kk"},
		{2, "libin"},
	}

	file, _ := os.Create("user.csv")
	writer := csv.NewWriter(file)
	for _, user := range users {
		writer.Write([]string{strconv.Itoa(user.ID), user.Name})
	}
	writer.Flush()
}
