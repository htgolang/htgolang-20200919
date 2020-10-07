package main

import (
	"fmt"
)

func main() {
	slice00 := []int{108, 107, 105, 109, 103, 102}

//	// 01 bubblesort
//	for j := 0; j < len(slice00)-1; j++ {
//		fmt.Println("============================Round", j+1)
//		for i := 0; i < len(slice00)-1-j; i++ {
//			fmt.Println("compare:", slice00[i], slice00[i+1])
//
//			if slice00[i] > slice00[i+1] {
//				fmt.Println("exchage:---> ", slice00[i], slice00[i+1])
//				slice00[i], slice00[i+1] = slice00[i+1], slice00[i]
//			}
//			fmt.Println("result:", slice00)
//		}
//		fmt.Println("Round result:", slice00)
//	}

	// 02 bubblesort optimized
	for i:=len(slice00)-1;i>0;i--{
		fmt.Println("+++++++++++++++++++++++++++++Round", i)
		for j:=0;j<i;j++{
			if slice00[j]> slice00[i]{
				fmt.Println("exchage:---> ", slice00[j], slice00[i])
				slice00[j],slice00[i]=slice00[i],slice00[j]
				fmt.Println("result:", slice00)
			}
		}
		fmt.Println("Round result:", slice00)
	}
}
//https://www.runoob.com/w3cnote/bubble-sort.html
