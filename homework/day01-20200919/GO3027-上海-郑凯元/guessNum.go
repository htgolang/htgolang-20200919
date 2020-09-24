package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Playnum() (player_num int) {

	player_num, _ = fmt.Print("请输入1-100的数字: ")
	fmt.Scan(&player_num)

	if player_num < 1 || player_num > 100 {
		player_num, _ = fmt.Print("笨蛋！输入1-100的数字：")
		fmt.Scan(&player_num)
	}

	return player_num

}

func GuessNum() {
	rand.Seed(time.Now().Unix())
	real_num := rand.Intn(100)
	fmt.Println(real_num)

START:
	for i := 5; i >= 0; i-- {
		player_num := Playnum()

		if i == 0 {
			var choose string
			fmt.Print("要不要在来一次？y/n:")
			fmt.Scan(&choose)

			if choose == "y" {
				goto START
			} else {
				fmt.Println("太笨了，游戏结束！")
			}
			continue
		}

		if player_num < real_num {
			fmt.Printf("小了,还有 %d 次\n", i)
			continue
		} else if player_num > real_num {
			fmt.Printf("大了,还有 %d 次\n", i)
			continue
		} else {
			fmt.Println("恭喜你，猜对了！")
			break
		}
	}

}

func main() {
	GuessNum()
}
