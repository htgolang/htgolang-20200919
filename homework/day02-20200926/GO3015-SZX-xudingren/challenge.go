package main

import (
	"fmt"
)

//a. 对问题2中的切片使用插入排序 实现从小到大排序
//b. 生成n个随机数, 在已排序的切片中(插入排序或者冒泡排序)查找其索引并打印, 未找到打印-1

//插入排序
func insertSort(sli []int) []int {
	//思想：从第一个元素开始，默认认为已排序
	//取出下一个元素，在已排序的切片中，从后往前依次比较，如果小于前一个元素，则交换位置
	for i := 1; i < len(sli); i++ {
		for j := i; j > 0 && j <= i; j-- {
			if sli[j-1] > sli[j] {
				sli[j-1], sli[j] = sli[j], sli[j-1]
			}
		}
	}
	return sli
}

//插入排序优化
func insertSort2(sli []int) []int {
	//思想：从第一个元素开始，默认认为已排序
	//取出下一个元素，在已排序的数组中，从后往前依次比较，如果小于前一个元素，则交换位置
	for i := 1; i < len(sli); i++ {
		for j := i; j > 0 && sli[j-1] > sli[j]; j-- {
			sli[j-1], sli[j] = sli[j], sli[j-1]
		}
	}
	return sli
}

//对已排序的切片，可以利用二分查找迅速查找到元素所在索引，时间复杂度O(logn)
func BinarySearch(sli []int, target int) int {
	low := 0
	high := len(sli) - 1
	for low < high {
		mid := (low + high) / 2
		if sli[mid] > target {
			high = mid - 1
		} else if sli[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	//插入排序
	sli := []int{108, 107, 105, 109, 103, 102}
	insertSort(sli)
	fmt.Println(sli)
	sli2 := []int{108, 107, 105, 109, 103, 102}
	insertSort2(sli2)
	fmt.Println(sli2)

	//生成随机数
	//randSli := []int{}
	//rand.Seed(time.Now().Unix())
	//for i:=0; i<10; i++ {
	//	randSli = append(randSli, rand.Intn(100))
	//}
	//二分查找
	sli3 := []int{56, 54, 37, 80, 63, 56, 14, 40, 5, 90}
	insertSort2(sli3)
	r := BinarySearch(sli3, 54)
	fmt.Println(sli3)
	fmt.Println("found it, index is:", r)
}
