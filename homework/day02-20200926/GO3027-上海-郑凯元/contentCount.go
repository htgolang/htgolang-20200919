package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func main() {
	inputFile := "MaGedu-Go/local_homework/day02/dream"
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "FIle Error: %s\n", err)
	}
	// 将从文件获取的内容格式从byte转换成string并转换成小写
	letterContent := strings.ToLower(string(content))

	letterCount := make(map[string]int)

	for _, v := range strings.ToLower(letterContent) {
		// 判断是否是字母
		if unicode.IsLetter(v) {
			letterCount[string(v)] = strings.Count(letterContent, string(v))
		}
	}
	fmt.Println(letterCount)

}
