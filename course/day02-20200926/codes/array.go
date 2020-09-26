package main

import "fmt"

func main() {
	// 定义变量names为元素类型为字符串长度为55的数组
	var names [55]string // string 55

	var scores = [...]int{100, 88}
	fmt.Printf("%T\n", names)
	fmt.Printf("%q\n", names)
	fmt.Printf("%T\n", scores)

	fmt.Println(names, scores)

	// 零值
	// n个对应元素类型的零值组成的数组

	// 索引 0 -> len-1
	// 0 1 2 3 4
	// 字面量
	// [length]type{v1, v2, ..., vn}
	// [length]type{i1:v1, ii:vi, in:vn}
	// [...]type{i1:v1, ii:vi, in:vn} // 最大的索引+1
	// [...]type{v1, v2, ..., vn} // 元素的数量

	// 操作
	// 关系运算 != ==
	var nums [2]int = [...]int{100, 88}
	fmt.Println(nums == scores)
	// 访问值 修改值
	// 索引 ()
	fmt.Println(nums[0])
	fmt.Println(nums[1])

	nums[0] = 101
	nums[1] = 102
	fmt.Println(nums)

	// 如何计算数组的长度
	fmt.Println(len(nums))

	// 遍历
	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	for v := range nums {
		fmt.Println(v, nums[v])
	}

	for i, v := range nums {
		fmt.Println(i, v)
	}

}
