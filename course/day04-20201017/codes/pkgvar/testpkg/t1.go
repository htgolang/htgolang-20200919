package testpkg

import "fmt"

// 包内成员可见性
// 首字母大写包外可见
// 首字母小写只能在包内可见

var T1Name = "T1"

var t3Name = "t3"

const T1Const = "T1"

func T1Func() {
	fmt.Println("T1 func")
	t3Func()
	fmt.Println(T1Name)
}

func t3Func() {
	fmt.Println("t3 func")
}
