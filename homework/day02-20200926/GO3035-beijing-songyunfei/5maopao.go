package main

import "fmt"

func main()  {
	arr := []int{108, 107, 105, 109, 103, 102}
	for true {
		flag := false
		for i:=0; i<len(arr)-1;i++{

			if arr[i] > arr[i+1]{
				arr[i],arr[i+1] = arr[i+1], arr[i]
				flag = true
			}
		}
		if flag{
			continue
		}else {
			break
		}
	}
	fmt.Println(arr)
}
