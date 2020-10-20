package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Unix())
	//模板 占位
	//年 => 2006
	//月 => 01
	//日 => 02
	//时 => 03(12h)/15(24h)
	//分 => 04
	//秒 => 05

	// 年-月-日 24时:分:秒
	format := "2006年01月02 15:04:05"
	fmt.Printf("%T\n", now.Format(format))
	fmt.Println(now.Format(format))

	// 2018-01-02 16:05 - 年-月-日 (24h)时:分 -> 2006-01-02 15:04

	// 前面时解析的格式， 后面时时间格式的字符串
	t1, err := time.Parse("2006-01-02 15:04", "2018/01-02 16:05")

	fmt.Println(err, t1)
}
