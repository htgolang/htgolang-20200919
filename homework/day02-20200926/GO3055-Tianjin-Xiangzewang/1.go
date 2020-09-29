package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// 读取文件，并转换成字符串
	b, _ := ioutil.ReadFile("/opt/gocode/homework/i_have_a_dream.txt")
	str_all := string(b)

	// 定义map
	dmp := make(map[string]int, 26)
	for i := 'a'; i <= 'z'; i++ {
		// 格式化字串
		x := fmt.Sprintf("%c", i)
		nums := strings.Count(str_all, x)
		dmp[x] = nums
	}
	fmt.Println(dmp)
}
