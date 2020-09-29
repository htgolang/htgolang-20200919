package main

import "fmt"

func main() {
	arr := []int{108, 107, 105, 109, 103, 102}
    for i := 1; i < len(arr); i++ {
        vi := arr[i]
        j := i
        for ; j > 0 && arr[j-1] >= vi; j-- {
            arr[j] = arr[j-1]
        }
        arr[j] = vi
    }
	fmt.Println(arr)
}


