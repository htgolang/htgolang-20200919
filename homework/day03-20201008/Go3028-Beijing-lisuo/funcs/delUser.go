package funcs

import (
	"fmt"
	_ "fmt"
	"strconv"
	"strings"
)

// one can use ID or name to find a user and del it
func DelUser() {
	var name string
	var input string
	fmt.Println("Who you want del(Id/Name)?")
	fmt.Scanln(&input)
	if s, err := strconv.Atoi(strings.TrimSpace(input)); err == nil {
		id := int64(s)
		fmt.Printf("idType: %T  idValue: %v", id, id)
	} else {
		name = strings.ToLower(strings.TrimSpace(input))
		fmt.Printf("nameType: %T  nameValue: %v\n", name, name)
	}
}
