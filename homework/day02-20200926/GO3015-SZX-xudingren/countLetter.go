package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func countLetterUseByMap() map[string]int {
	article, err := ioutil.ReadFile("/tmp/I_have_a_dream.txt")
	if err != nil {
		fmt.Printf("read file err, %s", err.Error())
	}
	lowerArticle := strings.ToLower(string(article))
	countLetter := make(map[string]int)
	for _, v := range lowerArticle {
		if unicode.IsLetter(v) {
			countLetter[string(v)]++
		}
	}
	return countLetter
}

//用于巩固学习，效率没第一个高
func countLetterUseByArray() map[string]int {
	article, err := ioutil.ReadFile("/tmp/I_have_a_dream.txt")
	if err != nil {
		fmt.Printf("read file err, %s", err.Error())
	}
	lowerArticle := strings.ToLower(string(article))
	countLetterArr := [26]int{}
	for _, v := range lowerArticle {
		if unicode.IsLetter(v) {
			countLetterArr[v-'a']++ // 利用字符可以进行数值运算的特性
		}
	}
	countLetterMap := make(map[string]int)
	for i, v := range countLetterArr {
		countLetterMap[string(byte(i)+'a')] = v
	}
	return countLetterMap
}

func main() {
	c1 := countLetterUseByMap()
	fmt.Println(c1)
	c2 := countLetterUseByArray()
	fmt.Println(c2)
}
