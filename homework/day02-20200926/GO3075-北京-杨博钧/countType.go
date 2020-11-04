package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//习题一统计I have a dream中各字母个数
func main() {
	//定义map记录结果
	var typeCounts = make(map[string]int)
	bytes, err := ioutil.ReadFile("file/I_have_a_dream.txt")
	if err != nil {
		fmt.Printf("打开文件错误:%v", err)
	}
	for _, v := range bytes {
		//对字母转小写
		value := []byte(strings.ToLower(string(v)))[0]
		//统计个数
		if value >= 97 && value <= 122 {
			typeCounts[string(value)] ++
		}
	}
	//打印结果
	fmt.Println(typeCounts)
}