package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var x string
	//生成个map
	mp := make(map[string]int, 26)
	for i := 'a'; i <= 'z'; i++ {
		x = fmt.Sprintf("%c", i)
		mp[x] = 0
	}
	bytes, err := ioutil.ReadFile("homework/day02-20200926/Go3060-北京-degary/files/I_have_a_dream.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	data := string(bytes)

	for i := 0; i < len(data); i++ {
		str := string(data[i])
		//转换小写
		str = strings.ToLower(str)
		//不存在 则跳过
		if _, ok := mp[str]; !ok {
			continue
		}
		mp[str] += 1
	}
	fmt.Println(mp)

}
