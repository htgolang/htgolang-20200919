package main

import "fmt"
func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			var sum int
			sum = i * j
			fmt.Printf("%v * %v = %v \t",j,i,sum)
			//fmt.Print("")
		}
		fmt.Println()
	}
}
