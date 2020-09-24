package main

import (
	"fmt"
	"math/rand"
)

func main() {
	max_gress := 6
	var answer int

	//num := rand.Intn(100)
	for i := 1; i < max_gress; i++ {
		num := rand.Intn(100)
		fmt.Println("请输入你的数字:")

		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Println("请重新输入")
			continue
		}

		if i < 5 {
			  if answer < num {
			  	fmt.Printf("第%v次 %v < %v 输入的数字小于随机数！\n\n", i , answer , num)
			  } else if answer > num {
				fmt.Printf("第%v次 %v > %v 输入的数字大于随机数！\n\n", i , answer , num)
			  } else {
				  fmt.Printf("恭喜成功！！\n 第%v次 %v = %v 输入的数字大于随机数！\n\n", i , answer , num)
				  break
			  }
		} else if i == 5 {
			if  answer == num {
				fmt.Printf("恭喜成功！！\n 第%v次 %v = %v 输入的数字大于随机数！\n\n", i , answer , num)

			} else {
				fmt.Printf("太笨了，程序结束!")
			}


		}



	}

}