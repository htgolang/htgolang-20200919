package main

import (
	"fmt"
	"pkgvar/testpkg"
)

func main() {
	testpkg.T1Func()
	testpkg.T2Func()
	// t3仅包内可见不能再外部调用
	// testpkg.t3Func()
	fmt.Println(testpkg.T1Name)
	testpkg.T1Name = "kkT1"
	fmt.Println(testpkg.T1Name)
	testpkg.T1Func()
	fmt.Println(testpkg.T1Const)
	// testpkg.T1Const = "kkConst"
}
