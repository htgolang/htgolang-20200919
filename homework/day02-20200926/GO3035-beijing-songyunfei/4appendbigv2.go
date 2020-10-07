package main

import (
	"fmt"
)

func main()  {
	arr :=  []int{108, 107, 105, 109, 103, 102}
	var big int
	var bigIndex int
	var secbig int
	var secbigindex int
	for i,v := range arr {
		if v > big {
			big = v
			bigIndex = i
		}
	}
	//和最后一位交换位置
	arr[bigIndex],arr[len(arr)-1] = arr[len(arr)-1], arr[bigIndex]
	fmt.Println(arr)
	//找出第二大的数字
	for sk,sv := range arr[:bigIndex]{
		if sv > secbig{
			secbig = sv
			secbigindex = sk
		}
	}
	//和倒数第二交换位置
	arr[secbigindex],arr[len(arr)-2] = arr[len(arr)-2],arr[secbigindex]
	fmt.Println(arr)


}

