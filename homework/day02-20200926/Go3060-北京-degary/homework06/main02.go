package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandInt(max, min int) (int, error) {
	var x int
	if max <= min {
		err := fmt.Errorf("%s\n", fmt.Sprintln("mix要大于min"))
		return 0, err

	}

	rand.Seed(time.Now().Unix())
	/*
		for {
			x = rand.Intn(n)
			if m < x && x < n {
				break
			}
		}

	*/
	x = rand.Intn(max - min)
	x += min
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
