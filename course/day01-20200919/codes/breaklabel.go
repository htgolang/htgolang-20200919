package main

import "fmt"

func main() {
END:
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			if i == 5 {
				break END
			}
			fmt.Println(j, i)
		}
	}
}
