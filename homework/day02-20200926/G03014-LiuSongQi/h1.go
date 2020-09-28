package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

/*
统计"我有一个梦想"英文文章中每个英文字母(不区分大小写字母)
(不统计标点符号)出现的次数 map for 关系运算 strings
*/

func readDreamFile(filename string) string {
	res, _ := ioutil.ReadFile(filename)
	return string(res)
}

func main() {
	content := readDreamFile("./dream.txt")
	var strCount = make(map[string]int)
	for _, v := range strings.ToLower(content) {
		isLower := unicode.IsLower(v)
		if isLower {
			strCount[string(v)]++
		}

	}
	fmt.Println(strCount)
}
