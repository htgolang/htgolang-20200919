package main

//4. 针对问题3, 将第二大的移动到切片的最后第二位

import (
	"fmt"
)

func main() {
	num_list := []int{108, 107, 105, 109, 103, 102}
	//var maxNum int

	for i := 0; i < len(num_list)-1; i++ {
		if num_list[i] > num_list[i+1] {
			num_list[i], num_list[i+1] = num_list[i+1], num_list[i]
		}

	}

	for i := 0; i < len(num_list)-1; i++ {
		if num_list[i] > num_list[i+1] {
			num_list[i], num_list[i+1] = num_list[i+1], num_list[i]
		}

	}

	fmt.Println(num_list)

}
