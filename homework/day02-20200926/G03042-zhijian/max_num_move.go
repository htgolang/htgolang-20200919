package main

/*
3. 针对问题2的切片, 将最大的数字移动到切片的最后一位
	原来的数字都在移动后的切片中都存在, 只是最大的数字再最后一位
*/

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
	fmt.Println(num_list)

}
