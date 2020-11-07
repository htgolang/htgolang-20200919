package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	path, _ := filepath.Abs("main.go")
	// 打印文件绝对路径
	// fmt.Println(filepath.Abs(path))

	// 返回path中最后一个元素
	fmt.Println(filepath.Base("/a/b/c/d/1.txt"))

	fmt.Println(filepath.Clean("/a/b/c/d/1.txt"))

	// 去掉path中的最后一个元素 [\a\b\c\d]
	fmt.Println(filepath.Dir("/a/b/c/d/1.txt"))

	// 返回path文件扩展名 [.txt]
	fmt.Println(filepath.Ext("/a/b/c/d/1.txt"))

	// 返回用分隔符替换路径中的每个斜杠('/')字符的结果。多个斜杠被多个分隔符替代。
	fmt.Println(filepath.FromSlash("/a/b/c/d/1.txt"))

	// Glob返回所有匹配模式的文件的名称，如果没有匹配的文件，则返回nil
	fmt.Println(filepath.Glob("*.txt"))

	// 判断mathfile下是否包含path
	fmt.Println(filepath.HasPrefix(path, "mathfile"))

	// 判断是否为绝对路径
	fmt.Println(filepath.IsAbs("D:/Golang-03/Go-day05/homework/mathFile/mathFile "))

	// 路径拼接
	fmt.Println(filepath.Join("a", "b", "c.txt"))

	//文件模式匹配
	fmt.Println(filepath.Match("*.txt", "1.txt"))

	// 将文件名和目录分离开
	fmt.Println(filepath.Split("/a/b/c/d/1.txt"))

	// 返回路径切片
	fmt.Println(filepath.SplitList("/a/b/c/d/1.txt"))

	// 递归返回路径下的文件
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
