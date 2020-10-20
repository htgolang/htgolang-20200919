package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //设置随机种子（设置一次即可）
	suiJi := int(rand.Intn(100))
	//fmt.Println("随机数是：",suiJi)
	//fmt.Println(suiJi)
	var cai int
	fmt.Print("请输入你猜测的数字: ")
	var ci = 1
	for ci <= 5 {
		fmt.Scan(&cai)

		switch {
		case suiJi < cai:
			fmt.Println("猜的太大了。还想猜的话继续输入猜的数字")
		case suiJi > cai:
			fmt.Println("猜的太小了。还想猜的话继续输入猜的数字")
		case suiJi == cai:
			fmt.Println("牛逼，你猜对了！ ")
		}
		ci++
		if ci > 5 {
			fmt.Println("太笨了，竟然猜5次没猜对")
		}
	}
	fmt.Println("随机数是：", suiJi)

}
