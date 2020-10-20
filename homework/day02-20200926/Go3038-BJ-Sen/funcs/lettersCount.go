package funcs

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func ReadFile(file string) string {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Bytes read:", len(bytes))
	//str := strings.ToUpper(string(bytes))
	////fmt.Println("String read:", string(bytes))
	//fmt.Println("String read:", str)
	//for i := 'A'; i <= 'Z'; i++ {
	//	//a := strings.ToLower(string(i))
	//	a := string(i)
	//	n := strings.Count(string(str), a)
	//	if n != 0 {
	//		fmt.Println(a, "出现了", n, "次")
	//	}
	//}
	return string(content)
}

func LettersCount() {
	content := ReadFile("day02/I have a dream")
	letters := strings.ToUpper(content)
	for i := 'A'; i <= 'Z'; i++ {
		//a := strings.ToLower(string(i))
		letter := string(i)
		n := strings.Count(string(letters), letter)
		if n != 0 {
			fmt.Println(letter, "出现了", n, "次")
		}
	}
}

func LettersCount2() {
	content := ReadFile("I have a dream")
	letters := strings.ToUpper(content)
	var letterCount = make(map[string]int)
	for _, v := range letters {
		//fmt.Println(v)
		isUpper := unicode.IsUpper(v)
		if isUpper {
			letterCount[string(v)]++
		}
	}
	for let, count := range letterCount {
		println(let, "出现次数为", count)
	}
	//fmt.Println(letterCount)
}