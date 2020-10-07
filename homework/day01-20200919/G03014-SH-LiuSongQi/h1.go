package main

import (
	"fmt"
)

/*
1*1=1
2*1=2 2*2=4
3*1=3 3*2=6 3*3=9
...
*/

func main() {
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d*%d=%d\t", x, y, x*y)
		}
		fmt.Println("")
	}

}
