package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random() int {
	rand.Seed(time.Now().Unix())
	seed := rand.Intn(20)
	// fmt.Println(seed)
	return seed
}

func main() {

	/*

		猜数字 生成随机数字0-100 从控制台数据 与生成数字比较 大 提示太大了 小 提示太小了 等于 成功, 程序结束

		最多猜测五次，未猜对，说太笨了，程序结束

	*/
	var number int
	i := 0
	sum := 1
	Endcount := 5

END:

	for i < Endcount {
		fmt.Print("请输入数字： ")
		fmt.Scan(&number)

		switch {
		case number > random():
			fmt.Printf("结果:[太大了], 重试次数:[%d]   数据对比：[输入值:[%d]  随机值:[%d]]\n", Endcount-i-1, number, random())

		case number < random():
			fmt.Printf("结果:[太小了], 重试次数:[%d]  数据对比：[输入值:[%d]  随机值:[%d]]\n", Endcount-i-1, number, random())

		default:
			fmt.Printf("结果:[答对了], 重试次数:[%d]  数据对比：[输入值:[%d]  随机值:[%d]]\n", Endcount-i-1, number, random())
			break END
		}
		i += sum

		if i >= 5 {
			fmt.Println("")
			fmt.Printf("您已失败:[%d]次，程序关闭.", Endcount)
		}
	}

}
