package main

import "fmt"

func main() {
	sliOri := []int{108, 107, 105, 109, 103, 102, 101}
	for j := len(sliOri); j > 0; j-- {
		for i := 0; i < j-1; i++ {
			if sliOri[i] > sliOri[i+1] {
				sliOri[i], sliOri[i+1] = sliOri[i+1], sliOri[i]
			}
		}
	}
	fmt.Println(sliOri)
}
