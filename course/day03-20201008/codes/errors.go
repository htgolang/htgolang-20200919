package main

import (
	"errors"
	"fmt"
	"strconv"
)

// error => 接口类型
func div(left, right int) (int, error) {
	if right == 0 {
		return 0, fmt.Errorf("right num is zero") //初始化error类型变量
	}
	return left / right, nil
}

func divV2(left, right int) (int, error) {
	if right == 0 {
		return 0, errors.New("right num is zero") //初始化error类型变量
	}
	return left / right, nil
}

func main() {
	// 需要给调用者返回错误信息
	// 通过函数最后一个返回值返回错误信息
	// 调用者需要对错误进行检查, 决定自己如何操作
	// 建议程序员去处理错误
	v, err := strconv.Atoi("123abc")
	if err == nil {
		fmt.Println("value是", v)
	} else {
		fmt.Println(err)
	}

	// 错误
	// 运行时错误
	// 		可恢复的错误 => 重试/忽略
	// 		不可恢复的错误 => 程序退出

	if v, err := div(1, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println("value is error", err)
	}
	fmt.Println(div(2, 1))
	fmt.Println(divV2(1, 0))
	fmt.Println(divV2(2, 1))
}
