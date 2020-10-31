package main

import (
	"encoding/csv"
	"fmt"
	"io"
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

	file, _ := os.Open("user.csv")
	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		id, _ := strconv.Atoi(line[0])

		users = append(users, User{id, line[1]})
	}
	fmt.Println(users)
}
