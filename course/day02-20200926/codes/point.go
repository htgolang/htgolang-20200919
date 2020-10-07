package main

import "fmt"

func main() {
	// 指针
	// 值类型
	// 赋值 原有的数据复制一份给新的变量
	// 两个变量之间没有任何联系
	// 对nums进行任何修改都不会影响nums2
	// 对nums2进行任何修改也不会影响nums

	// 有联系 => 引用类型(内部使用使用了指针/地址)

	// 数组值类型
	nums := [5]int{1, 2, 3, 4, 5}
	nums2 := nums

	fmt.Println(nums, nums2)

	nums2[0] = 100
	fmt.Println(nums, nums2)
	// int, float, bool, string => 值类型

	// var name *type

	var age = 1
	var agePointer *int = nil // 定义int类型的指针

	agePointer = &age // 取出age的地址赋值给agePointer => 取引用

	fmt.Println(agePointer, age)

	fmt.Println(*agePointer) // 获取agePointer内存地址中对应的存储的值 => 解引用

	*agePointer = 33
	fmt.Println(age, *agePointer)

	var numsPoint *[5]int = &nums
	fmt.Printf("%T\n", numsPoint)
	numsPoint[0] = 100
	fmt.Println(nums, numsPoint)

}
