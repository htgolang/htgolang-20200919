package main

//5. 使用冒泡排序算法对问题2切片进行从小到大排序

import (
	"fmt"
)

func main() {
	num_list := []int{108, 107, 105, 109, 103, 102}
	//var maxNum int
	for n := 0; n <= len(num_list); n++ {

		for i := 0; i < len(num_list)-1; i++ {
			if num_list[i] > num_list[i+1] {
				num_list[i], num_list[i+1] = num_list[i+1], num_list[i]
			}

		}
	}

	fmt.Println(num_list)

}
