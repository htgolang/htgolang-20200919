package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	var tmp int
	var p int
	rand.Seed(time.Now().Unix())
	guess := rand.Intn(100)
	//fmt.Println(guess)
	limit := 5
	i:=1
	for i <= limit{
		fmt.Printf("请输入数字(0-100):")
		_,_ = fmt.Scan(&tmp)
		s := limit-i
		if tmp < guess{
			fmt.Printf("小了...还有%d次机会.\n", s)
			i++
		}else if tmp == guess{
			fmt.Println("恭喜你...")
			p = 1
			break
		}else {
			fmt.Printf("大了...还有%d次机会.\n", s)
			i++
		}
	}
	if p == 1 {
		fmt.Println("你很棒!!!")
	}else {
		fmt.Println("五次机会用完了没有猜对太笨了...")
	}


}
