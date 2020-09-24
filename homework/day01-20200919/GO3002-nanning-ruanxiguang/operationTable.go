// 写法一
package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}

/* 写法二

package main

import "fmt"

func main() {
	for i := 1; i < 10; i++{
		var data string
		for j := 1; j <= i; j++{
			if j < 2{
				data += fmt.Sprintf("%vX%v=%-2d", i, j, i*j)
			}else {
				data += fmt.Sprintf("%vX%v=%-3d", i, j, i*j)
			}
		}
	fmt.Println(data)
	}
}

*/
