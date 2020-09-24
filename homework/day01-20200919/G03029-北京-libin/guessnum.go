package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	gameRound  int = 5
	compareNum int
	yourChoice string
)

func main() {
	rand.Seed(time.Now().UnixNano())
GAME:
	fmt.Println("剩下的游戏次数是", gameRound)
	for gameRound > 0 {

		randNum := rand.Intn(100)
		fmt.Scanln(&compareNum)
		if randNum == compareNum {
			fmt.Println("bingo!")
			goto FIN
		} else if randNum > compareNum {
			fmt.Println("随机生成的数字比较大，请继续输入数字")
			// continue
		} else {
			fmt.Println("手动输入的数字比较大，请继续输入数字")
		}
		gameRound--
		fmt.Printf("还剩下%d次数\n", gameRound)
	}

FIN:
	fmt.Println("是否继续猜数字游戏，如果继续，请输入 y")
	fmt.Scanln(&yourChoice)
	if yourChoice != "y" {

		fmt.Println("游戏结束，再见")
		os.Exit(0)
	} else {
		gameRound = 5
		goto GAME
	}

}
