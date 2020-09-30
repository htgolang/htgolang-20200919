package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func main()  {
	cont := make(map[string]int)
	f,err:= ioutil.ReadFile("./doc.txt")
	if err != nil{
		fmt.Println(err)
	}
	for _,v := range strings.ToUpper(string(f)){
		if unicode.IsUpper(v){
			cont[string(v)]++
		}

	}
	for s,c := range cont{
		fmt.Printf("字母:%s 出现次数:%d\n",s,c)
	}
}
