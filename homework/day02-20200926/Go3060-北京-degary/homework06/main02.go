package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandInt(m, n int) (int, error) {
	var x int
	if n <= m {
		err := fmt.Errorf("%s\n", fmt.Sprintln("第二个参数要大于第一个参数"))
		return 0, err

	}
	rand.Seed(time.Now().Unix())
	x = rand.Intn(n - m)
	x += m
	return x, nil
}

func main() {
	sliOri := []int{101, 102, 103, 104, 105, 106, 107, 108}
	num, err := getRandInt(100, 120)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("The random num is %d\n", num)
	for i := 0; i < len(sliOri); i++ {
		if num == sliOri[i] {
			fmt.Printf("key: %d,value: %d", i, sliOri[i])
			goto END
		}
	}
	fmt.Println(-1)
END:
}
