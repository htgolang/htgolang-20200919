package main

import "fmt"

func main() {
	for r := 1; r <= 9; r++ { //定义行的行数定义，1-9行.
		for k := 1; k <= r; k++ {
			fmt.Printf("%d * %d = %d\t", k, r, k*r)
		}
		fmt.Println("\n") //外循环，每打印完一行，进行一次换行;
	}
}

#########===============================================
package main

import "fmt"

func main() {
        for s := 9; s >= 1; s-- {
                for d := 1; d <= s; d++ {
                        fmt.Printf("%d * %d = %d\t", d, s, s*d)
                }
                fmt.Println("\n")
        }
}
