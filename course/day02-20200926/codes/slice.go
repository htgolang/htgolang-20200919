package main

import "fmt"

func main() {
	// 定义names 元素类型为string的切片
	var names []string // nil

	fmt.Printf("%T\n", names)
	fmt.Printf("%q\n", names)

	names = []string{"赵昌建", "sen"}
	fmt.Printf("%T\n", names)
	fmt.Printf("%q\n", names)

	// 字面量
	// []type{} => 空切片
	// []type{v1, v2, ..., vn}
	// []type{i1:v1, i2:v2, in:vn}
	names = []string{1: "赵昌建", 10: "sen"}
	fmt.Printf("%T\n", names)
	fmt.Printf("%q\n", names)

	// 访问 修改元素
	// 索引
	fmt.Println(names[1])
	fmt.Println(names[10])
	fmt.Println(names[9])
	names[9] = "卫智鹏"
	fmt.Println(names)

	fmt.Println(len(names))
	// 长度 切片中已经存在元素的数量
	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	for v := range names {
		fmt.Println(v, names[v])
	}

	for i, v := range names {
		fmt.Println(i, v)
	}

	// 添加元素
	// 切片的末尾
	names = append(names, "刘冉")
	fmt.Printf("%q\n", names)
	fmt.Println(len(names))

	// 删除元素
	// 切片操作
	// names[start:end] names中从start开始到end-1所有元素组成的切片
	// names[1:10]
	//			0	1								9	 10		11
	// names = ["" "赵昌建" "" "" "" "" "" "" "" "卫智鹏" "sen" "刘冉"]
	// [names[1], names[2], ... names[9]]
	fmt.Printf("%q\n", names[1:10])
	// 删除索引为0 如果索引为len - 1 元素
	// 0, len(names)-1
	names = names[1:len(names)]
	fmt.Printf("%q\n", names)
	names = names[0 : len(names)-1]
	fmt.Printf("%q\n", names)
	// 删除中间的元素
	nums := []int{0, 1, 2, 3, 4, 5}
	// 删除3
	// copy(dst, src) src->dst
	nums2 := []int{10, 11, 12, 13, 14, 15, 16}

	// nums2多
	copy(nums, nums2)
	fmt.Println(nums, nums2)
	// nums[0:3], nums[4:5]
	copy(nums[3:len(nums)], nums[4:len(nums)])
	// 4, 5, 5			4, 5
	// nums 不应该变
	fmt.Println(nums[0 : len(nums)-1])
	// 切片操作和原来的切片共享存储空间

}
