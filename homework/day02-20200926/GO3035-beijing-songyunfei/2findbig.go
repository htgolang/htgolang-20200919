package main

import "fmt"

func main()  {
	arr :=  []int{108, 107, 105, 109, 103, 102}
	var big int
	for i:=0;i<len(arr);i++{
		if arr[i] > big{
			big=arr[i]
		}
	}
	fmt.Println("Big:",big)
}
