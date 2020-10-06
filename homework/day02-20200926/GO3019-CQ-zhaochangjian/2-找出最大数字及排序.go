package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 2. int切片 []int{108, 107, 105, 109, 103, 102} 找出最大的数字
	nums := []int{108, 107, 105, 109, 103, 102}
	// 假设nums[0]为最大值
	maxNum := nums[0]
	maxIndex := 0 // 最大值索引
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			maxIndex = i
		}
	}
	// 打印最大值
	fmt.Println(maxIndex, maxNum)

	// 3. 针对问题2的切片, 将最大的数字移动到切片的最后一位 原来的数字都在移动后的切片中都存在, 只是最大的数字再最后一位
	dst := nums[maxIndex:len(nums)]
	src := nums[maxIndex+1 : len(nums)]
	copy(dst, src) // 删除最大值
	nums = nums[0 : len(nums)-1]
	nums = append(nums, maxNum) // 末尾追加最大值
	fmt.Println(nums)

	// 4. 针对问题3, 将第二大的移动到切的最后第二位
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
	fmt.Println(nums)

	// 5. 使用冒泡排序算法对问题2切片进行从小到大排序
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)

	// 6. 生成n个随机数, 在已排序的切片中(插入排序或者冒泡排序)查找其索引并打印, 未找到打印-1
	rand.Seed(time.Now().Unix()) // 随机数种子
	for j := 0; j <= 10; j++ {
		var n int = rand.Intn(100) + 100 // 生成100-200随机数
		flag := false
		for i, v := range nums {
			if v == n {
				fmt.Println(i, v)
				flag = true
			}
		}
		if !flag {
			fmt.Println(-1)
		}

	}

}
