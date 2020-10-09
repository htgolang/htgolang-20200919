package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNums(n int) []int {
	newSli := []int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		randNums := rand.Intn(100) + 100
		newSli = append(newSli, randNums)
	}
	// fmt.Println(newSli)
	return newSli
}

func insertionSort(old []int) []int {
	n := len(old)

	for i := 1; i < n; i++ {
		for j := i; j > 0 && old[j] < old[j-1]; j-- {
			old[j], old[j-1] = old[j-1], old[j]
		}
	}
	return old
}

func main() {
	nums := []int{108, 107, 105, 109, 103, 102}
	nums1 := insertionSort(nums)
	fmt.Println(nums1)

	newSli := randNums(1)
	newSli1 := insertionSort(newSli)
	fmt.Println(newSli1)

	for k, v := range nums {
		for _, v1 := range newSli {
			if v == v1 {
				fmt.Println("%c, %v\n", k, v)
			} else {
				fmt.Println("-1")
			}
		}

	}
}
