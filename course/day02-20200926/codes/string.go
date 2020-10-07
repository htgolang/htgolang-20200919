package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string = "kk"
	var desc string = "i love china"
	// 字面量 "" ``
	// 零值 ""
	// 操作
	// 连接 +
	// 关系运算 == != > < >= <=
	// 赋值操作 = +=
	// 长度 len 字节长度
	// 索引 字节
	// 切片 => 字符串 字节
	fmt.Println(name, desc)
	fmt.Println(desc[1:5])
	var txt = "我爱中国"
	for i, v := range txt {
		fmt.Printf("%d, %q\n", i, v)
	}

	fmt.Println(utf8.RuneCountInString(txt))
	fmt.Println(utf8.RuneCountInString(desc))

	// 字符串 => []byte
	fmt.Println([]byte(desc))
	fmt.Println(string([]byte(desc)))
	fmt.Println([]byte(txt))
	fmt.Println(string([]byte(txt)))

	// 字符串 => []rune
	fmt.Println([]rune(desc))
	fmt.Println(string([]rune(desc)))
	fmt.Println([]rune(txt))
	fmt.Println(string([]rune(txt)))
}
