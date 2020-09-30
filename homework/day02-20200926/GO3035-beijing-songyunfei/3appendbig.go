package main

import "fmt"

func main()  {
	arr :=  []int{108, 107, 105, 109, 103, 102}
	var big int
	var bigIndex int
	for i,v := range arr{
		if v > big{
			big=v
			bigIndex=i
		}
	}
	if bigIndex == 0 {
		newarr := append(arr[1:],big)
		fmt.Println(newarr)
	}else if bigIndex == len(arr){
		fmt.Println(arr)
	}else {
		newarr := append(arr[:bigIndex],arr[bigIndex+1:]...)
		newarr = append(newarr,big)
		fmt.Println(newarr)
	}
}
