package main

import "fmt"

func main() {
	//[length]type
	// type [2]int
	// 二维数组
	var ms [3][2]int

	fmt.Printf("%T\n", ms)
	fmt.Println(ms)

	fmt.Printf("%T %v\n", ms[0], ms[0])
	fmt.Printf("%T %v\n", ms[0][0], ms[0][0])

	ms = [...][2]int{
		1: [2]int{1, 2},
		2: [2]int{3, 4},
		0: [2]int{5, 6},
	}
	// [[5, 6], [1, 2], [3, 4]]

	ms[0][1] = 100
	fmt.Println(ms)

	ms[1] = [2]int{101, 102}
	fmt.Println(ms)

	for i, line := range ms {
		fmt.Println(i, line)
		for j, v := range line {
			fmt.Println(i, j, v)
		}
	}
}
