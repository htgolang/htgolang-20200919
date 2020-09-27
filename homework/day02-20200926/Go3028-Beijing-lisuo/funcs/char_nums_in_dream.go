package funcs

import (
	"fmt"
	"io/ioutil"
)

func charNumsInDream() {
	data, err := ioutil.ReadFile("../files/I_have_a_dream.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Println(string(data))
}
