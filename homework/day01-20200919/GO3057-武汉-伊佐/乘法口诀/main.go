/*
打印乘法口诀
*/

package main

import "fmt"

// miltiplicationTables
func miltiplicationTables(num int) {
	// 打印乘法口诀
	for i := 1; i <= num; i++ {
		for s := 1; s <= i; s++ {
			fmt.Printf("%d * %d = %d \t", s, i, i*s)
		}
		fmt.Println()
	}
}

func main() {
	num := 9
	miltiplicationTables(num)
}
