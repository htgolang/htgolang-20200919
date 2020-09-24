package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomnum := rand.Intn(100)
	fmt.Println("请输入数字: ")
	cycles := 5
	for cycles > 0 {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("$ ")
		if input, err := inputReader.ReadString('\n'); err == nil {
			inputnum := strings.Trim(input, " ")
			inputnum = strings.TrimSpace(inputnum)
			finalnum, err := strconv.Atoi(inputnum)
			if err != nil {
				fmt.Println("你输入的不是数字,请重新输入...")
				continue
			}
			switch {
			case finalnum > randomnum:
				fmt.Println("你输入数字太大了😌")
			case finalnum < randomnum:
				fmt.Println("你输入数字太小了😌")
			case finalnum == randomnum:
				fmt.Println("你猜对了👏")
			}
			if cycles == 1 {
				fmt.Println("你太笨了,再见😏")
			}
			cycles--
		}
	}
}