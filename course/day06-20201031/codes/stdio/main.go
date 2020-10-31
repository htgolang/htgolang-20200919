package main

import "fmt"

func main() {
	// 标准输入 => 命令行的文件 os.Stdin
	// 标准输出 => 命令行的文件 os.Stdout
	// 标准错误输出 => 命令行的文件 os.Stderr
	// content := make([]byte, 3)

	// fmt.Print("请输入内容：")
	// fmt.Println(os.Stdin.Read(content))

	// //fmt.Scan
	// fmt.Printf("%q\n", string(content))

	// os.Stdout.WriteString("我是Stdout的输出")
	// fmt.Print("xxxx")

	// // fmt.Print / Println / Printf
	// // fmt.Sprint / Sprintln / Sprintf
	// // fmt.Fxxx
	// fmt.Fprint(os.Stdout, "aaaaa")
	// fmt.Fprintln(os.Stdout, "aaaaa")
	// fmt.Fprintf(os.Stdout, "i am: %s", "aaaaa")

	// fmt.Scan => Scanln, Scanf
	// fmt.Sscan => 从字符串扫描到变量
	// fmt.Fscan => 从文件扫描到变量

	name := ""
	// fmt.Println(fmt.Scan(&name))
	// fmt.Printf("%q\n", name)

	// fmt.Println(fmt.Scanln(&name))
	// fmt.Printf("%q\n", name)
	fmt.Println(fmt.Scanf("name: %s", &name))
	fmt.Printf("%q\n", name)

}
