package main

import "fmt"

func main() {
	sli := []int{108, 107, 105, 109, 103, 102}
	tmp := sli[0]
	for i := 1; i < len(sli); i++ {
		if sli[i] > tmp {
			tmp = sli[i]
		}
	}
	fmt.Printf("最大数字是%v\n", tmp)
}
