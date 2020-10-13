package main

import (
	"fmt"

	lisuo "github.com/htgolang/htgolang-20200919/tree/master/homework/day02-20200926/Go3028-Beijing-lisuo/funcs"
)

func main() {
	var numsArray = []int{108, 107, 105, 109, 103, 102}
	fmt.Println("Running func CharNumsInDream...")
	lisuo.CharNumsInDream()

	fmt.Println("\nRunning func GetMaxNum...")
	var numsArrayMax = []int{108, 107, 105, 109, 103, 102}
	max, idx := lisuo.GetMaxNum(numsArrayMax)
	fmt.Printf("The max number is: %v , index is: %v\n", max, idx)

	fmt.Println("\nRunning func MoveToLast...")
	movedNums := lisuo.MoveToLast(numsArrayMax)
	fmt.Printf("Move max to last: %v\n", movedNums)

	fmt.Println("\nRunning func BubbleSort...")
	fmt.Printf("The original  numsArray: %v\n", numsArray)
	lisuo.BubbleSort(numsArray)
	fmt.Printf("Sorted numsArray: %v\n", numsArray)
}
