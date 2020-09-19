package main

import (
	"fmt"
	"unsafe"
)

func main() {
	isBody := false
	fmt.Println(isBody)

	fmt.Printf("%T %t\n", isBody, isBody)
	fmt.Println(unsafe.Sizeof(isBody))
	// 操作
	// 逻辑运算
	// and, or, not
	// 与操作 两个操作数 左操作数 && 右操作数
	// 运算规则: 都为真的结果为真
	// A && B   	A(true)	   A(false)
	//  B(true)		true		false
	//  B(false)	false		false
	fmt.Println("&&")
	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	// 或运算, 左操作数 || 右操作数
	// 运算规则: 只要有一个为真结果为真
	// A || B   	A(true)	   A(false)
	//  B(true)		true		true
	//  B(false)	true		false
	fmt.Println("||")
	fmt.Println(true || true)
	fmt.Println(false || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	// 非, !操作数
	// 运算规则：你真我假，你假我真
	// !A  A(true) A(false)
	//		false   true
	fmt.Println("!")
	fmt.Println(!true)
	fmt.Println(!false)

	fmt.Println("==")
	fmt.Println(true == true)
	fmt.Println(true == false)
	fmt.Println(false == false)
	fmt.Println(false == true)

	fmt.Println("!=")
	fmt.Println(true != true)
	fmt.Println(true != false)
	fmt.Println(false != false)
	fmt.Println(false != true)
}
