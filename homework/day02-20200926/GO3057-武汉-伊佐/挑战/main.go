package main

import (
	"fmt"
)

func main() {
	data := []int{108, 107, 105, 109, 103, 102}
	fmt.Printf("%v \n", data)

	// 插入排序
	// 思想:  假设现在检查的值左边的都是有序集合，做排序是只检查左边相邻一个值的大小，如果左边的值小于此次的值，那么不做检查循环直接通过
	//       如果左边相邻的值大于此次循环的值，那么对左边的序列做一次循环检查，找到比当前值大的值替换位置。
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			for num := 0; num < i; num++ {
				if data[num] > data[i] {
					data[num], data[i] = data[i], data[num]
				}
				// fmt.Println(data) // 不熟悉的可以打开这个打印查看对应的迭代变化
			}
		}

	}
	fmt.Println(data)

}
