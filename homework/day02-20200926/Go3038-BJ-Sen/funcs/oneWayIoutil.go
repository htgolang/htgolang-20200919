package funcs

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func OneWay() {
	bytes, err := ioutil.ReadFile("day02/I have a dream")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bytes read:", len(bytes))
	str := strings.ToUpper(string(bytes))
	//fmt.Println("String read:", string(bytes))
	fmt.Println("String read:", str)

	for i := 'A'; i <= 'Z'; i++ {
		//a := strings.ToLower(string(i))
		a := string(i)
		n := strings.Count(string(str), a)
		if n != 0 {
			fmt.Println(a, "出现了", n, "次")
		}
	}
}