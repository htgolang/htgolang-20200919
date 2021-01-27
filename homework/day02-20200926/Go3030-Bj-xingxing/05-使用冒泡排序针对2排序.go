package main

import (
	"fmt"
)

func main() {
	num := []int{108,107,105,109,103,102}

	for  j :=0; j<len(num)-1; j++{
	for i := 0; i <len(num)-1; i++ {
		if num[i] > num[i+1]{
			num[i],num[i+1] = num[i+1],num[i]
		}
	}
	}
	fmt.Println(num)

}