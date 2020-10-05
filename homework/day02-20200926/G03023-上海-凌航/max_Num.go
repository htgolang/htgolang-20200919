package main

import "fmt"

func main() {

	nums := []int{108, 107, 105, 109, 103, 102} //
	fmt.Println(nums[0], nums[1], nums[2], "##########")

	// 2、3 找出最大的数字，将最大的数字移动到切片的最后一位

	var maxNum int
	for k, v := range nums {
		if v > maxNum {
			maxNum = v
			temp := nums[k] // 最大的数据
			nums[k] = nums[len(nums)-1]
			nums[len(nums)-1] = temp
		} else if v < maxNum {

		}

	}
	fmt.Println("最大的数字为", maxNum)
	fmt.Println("将最大的数字移动到最后的切片为", nums)
	// 输出为
	//	最大的数字为 109
	//	将最大的数字移动到最后的切片为 [102 107 105 108 103 109]

	// 4 将第二大的移动到切片的最后第二位
	maxNum2 := nums[0]
	secodNum := 0
	for i := 0; i < len(nums)-1; i++ {
		if maxNum2 < nums[i+1] {
			secodNum = maxNum2
			maxNum2 = nums[i+1]
		} else {
			if secodNum < maxNum2 {
				if nums[i] > secodNum {
					secodNum = nums[i]
					temp2 := nums[i] // 第二大的数据为
					nums[i] = nums[len(nums)-2]
					nums[len(nums)-2] = temp2
				}
			}
		}

	}

	fmt.Println("将第二大的移动到切片的最后第二位为", nums)
	// 输出为
	//将第二大的移动到切片的最后第二位为 [102 103 105 107 108 109]


	//5 使用冒泡排序算法对问题2切片进行从小到大排序
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			//fmt.Println(nums[j],"==============")
			if nums[i] > nums[j] { // 前面的值大于后面的值 进行交换
				temp := nums[i]      // 将nums[i] 的值赋予给temp，此时temp 为nums[i] 的初始值比较为2个数中的较大的值
				nums[i] = nums[j] // 将nums[j]的值 赋予给nums[i],此时nums[i] 为2个数中的较小值
				nums[j] = temp  // 将 temp(较大的值) 赋予给nums[j]，交换完成
			}
		}
	}

	fmt.Println("使用冒泡排序后的结果为...",nums)
	// 输出
		//使用冒泡排序后的结果为... [102 103 105 107 108 109]



}
