package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 设置随机数种子(设置一次)

	fmt.Println(rand.Intn(100))
}
