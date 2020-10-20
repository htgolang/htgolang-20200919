package main

// 包名为testpkg
// 目录名称pkg

import (
	"testpkgname/pkg" // 导入目录
)

// 包名与文件名不一致
func main() {
	testpkg.Test() // 包名调用
}
