package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
b. 生成n个随机数, 在已排序的切片中(插入排序或者冒泡排序)查找其索引并打印, 未找到打印-1
*/

func randNums() []int {
	var nums = []int{}
	rand.Seed(time.Now().Unix()) //设置随机数种子
	for i := 0; i < 10; i++ {
		var randNum int = rand.Intn(200)
		nums = append(nums, randNum)
	}
	return nums
}

func sortNums(nums []int) []int {
	for x := 1; x < len(nums); x++ {
		k := nums[x]
		i := x - 1

		for i >= 0 && nums[i] > k {
			nums[i+1] = nums[i]
			i--

		}
		nums[i+1] = k
	}

	return nums
}

func main() {
	var nums = []int{108, 107, 105, 109, 103, 102}
	randNums := randNums()
	sortNums := sortNums(nums)

	fmt.Println(randNums, sortNums)
	for x, y := range randNums {
		for k, v := range sortNums {
			if y == v {
				fmt.Printf("随机数切片中的索引为: %d, 排序切片中的索引为: %d \n", x, k)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
