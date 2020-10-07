package main

import (
	"errors"
	"fmt"
)

//2. int切片 []int{108, 107, 105, 109, 103, 102}
//找出最大的数字
//
//3. 针对问题2的切片, 将最大的数字移动到切片的最后一位
//原来的数字都在移动后的切片中都存在, 只是最大的数字再最后一位
//4. 针对问题3, 将第二大的移动到切片的最后第二位
//使用冒泡排序算法对问题2切片进行从小到大排序
//
//提醒:
//int a, b = 1, 4
//交换方式一:
//int tmp = a
//a=b
//b = tmp
//交换方式二:
//a, b = b, a

// 判断切片中的最大值，时间复杂度O(n)
func MaxNum(sli []int) (int, error) {
	var max int
	if sli == nil {
		return max, errors.New("Null array")
	}
	for i := 0; i < len(sli); i++ {
		if i == 0 {
			max = sli[0]
			continue
		}
		if max >= sli[i] {
			continue
		}
		max = sli[i]
	}
	return max, nil
}

//最大值放到切片末尾，时间复杂度O(n)
func moveMaxToLast(sli []int) []int {
	// 省略合法校验
	var index int
	var max int
	for i := 0; i < len(sli); i++ {
		if i == 0 {
			index, max = i, sli[0]
			continue
		}
		if max >= sli[i] {
			continue
		}
		index, max = i, sli[i]
	}
	if index == len(sli)-1 {
		return sli
	} else {
		sli[index], sli[len(sli)-1] = sli[len(sli)-1], sli[index]
		return sli
	}
}

// 第二大值移动到倒数第二，时间复杂度O(n)
func moveSecondMaxToLast2(sli []int) []int {
	// 省略合法校验
	var index int
	var secondMax int
	for i := 0; i < len(sli)-2; i++ {
		if i == 0 {
			index, secondMax = i, sli[0]
			continue
		}
		if secondMax >= sli[i] {
			continue
		}
		index, secondMax = i, sli[i]
	}
	if index == len(sli)-2 {
		return sli
	} else {
		sli[index], sli[len(sli)-2] = sli[len(sli)-2], sli[index]
		return sli
	}
}

func bubbleSort(sli []int) []int {
	// 省略合法校验
	for i := 0; i < len(sli); i++ {
		for j := 1; j < len(sli)-i; j++ {
			if sli[j-1] > sli[j] {
				sli[j-1], sli[j] = sli[j], sli[j-1]
			}
		}
	}
	return sli
}

func main() {
	// 最大值
	sli1 := []int{108, 107, 105, 109, 103, 102}
	max, err := MaxNum(sli1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(max)
	// 最大值移动到末尾
	sli2 := []int{108, 107, 105, 109, 103, 102}
	moveMaxToLast(sli2)
	fmt.Println(sli2)

	// 第二大值移动到倒数第二
	moveSecondMaxToLast2(sli2)
	fmt.Println(sli2)

	// 冒泡排序
	sli3 := []int{108, 107, 105, 109, 103, 102}
	bubbleSort(sli3)
	fmt.Println(sli3)
}
