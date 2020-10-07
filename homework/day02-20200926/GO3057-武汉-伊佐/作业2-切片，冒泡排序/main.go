package main

import (
	"fmt"
)

func main() {
	data := []int{108, 107, 105, 109, 103, 102}
	fmt.Printf("源数据: %v\n", data)

	// 找到最大的数字，将它放到列表的最后一位
	for i, v := range data {
		if v >= data[len(data)-1] {
			data[i], data[len(data)-1] = data[len(data)-1], data[i]
		}
	}
	fmt.Printf("找出最大的数据放到最后一位: %v \n", data)

	// 找到第二大的数字，将它放到列表的倒数第二位
	for num := len(data) - 2; num < len(data); num++ {
		for i, v := range data {
			if v >= data[num] {
				data[i], data[num] = data[num], data[i]
			}
		}
	}
	fmt.Printf("找出最大及其第二大的数据，放置最后两位: %v \n", data)

	// 冒泡排序算法，从小到大排序
	for num := 0; num < len(data); num++ {
		for i, v := range data {
			if v >= data[num] {
				data[i], data[num] = data[num], data[i]
			}
		}
	}
	fmt.Printf("冒泡排序算法，从小到大排序: %v \n", data)

}
