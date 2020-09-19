package main

import "fmt"

func main() {
	var name string = `kk`

	fmt.Printf("%T, %s\n", name, name)

	fmt.Println("my name is " + name)

	// 关系运算
	// > >= < <= != ==
	// s1 s2
	// 字符串最左侧的字符开始比较
	// abc > abd
	// a b c > d false
	// abc > abcd false

	// =, +=
	// a += b // a = a + b

	name += " hi"
	fmt.Println(name)
	// name = "我爱中国"

	// 字符串内的数据=>ascii byte=>unicode n byte
	// 索引
	// name => "kk hi" 0->len(n)-1
	fmt.Printf("%T\n", name[0])
	// 长度 len()
	fmt.Println(len(name))
	// 切片
	fmt.Println(name[0:4])
}
