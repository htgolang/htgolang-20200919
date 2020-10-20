package main

import (
	"fmt"
	"sort"
)

func main() {

	/*

		2.int切片 []int{108, 107, 105, 109, 103, 102} 找出最大的数字

	*/
	var maxnumber []int

	maxnumber = []int{108, 107, 105, 109, 110, 120, 130, 103, 102}

	// sort进行又大到小排序
	sort.Ints(maxnumber)

	fmt.Println(maxnumber[len(maxnumber)-1])

}
