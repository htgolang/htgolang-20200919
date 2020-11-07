package main

import "fmt"

// 定义接口
/* type 接口名称 interface {
	方法声明(行为) => 签名
}
*/
type Writer interface {
	Write([]byte) (int, error)
}

type MyFile struct {
}

func (f *MyFile) Write() {

}

type MyFile2 struct {
}

func (f *MyFile2) Write(ctx []byte) (int, error) {
	return 0, nil
}

type MyFile3 struct {
}

func (f MyFile3) Write(ctx []byte) (int, error) {
	return 0, nil
}

// 自动生成指针接收者
/*
func (f *MyFile3) Write(ctx []byte) (int, error) {
	return 0, nil
}
*/

func main() {
	var writer Writer
	fmt.Printf("%T, %#v\n", writer, writer)

	// var myFile MyFile
	// writer = myFile
	// myFile 无Write方法

	// var myFile2 MyFile2
	// writer = myFile2
	// 值接收者无Write方法
	var myFile2Pointer = new(MyFile2)
	writer = myFile2Pointer
	fmt.Printf("%T, %#v\n", writer, writer)

	var myFile3 MyFile3
	writer = myFile3
	fmt.Printf("%T, %#v\n", writer, writer)

	var myFile3Pointer = &MyFile3{} //new(MyFile3)
	writer = myFile3Pointer
	fmt.Printf("%T, %#v\n", writer, writer)

	// 结构体 MyFile3 变量赋值
	// 值类型变量 myFile3 := MyFile3{}
	//			 var myFile3 MyFile
	// 指针类型变量 var myFile3Pointer *MyFile3
	// myFile3Pointer = &myFile3
	// myFile3Pointer = &MyFile3{}
	// myFile3Pointer = new(MyFile3) new函数 创建结构体的指针类型变量

}
