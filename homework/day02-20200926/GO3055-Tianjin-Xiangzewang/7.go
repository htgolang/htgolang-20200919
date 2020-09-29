package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{102, 103, 105, 107, 108, 109}

	rand.Seed(time.Now().Unix())

	rand_num := rand.Intn(200 - 100) + 100
	fmt.Println(rand_num)
	for k, v := range nums {
		if rand_num == v {
			fmt.Printf("元素%d的索引为:%d\n", v, k)
		} else {
			fmt.Println("-1")
		}
	}

}
