package main

import (
	"fmt"
)

func main() {
	sl := []int{108, 107, 105, 109, 103, 102}
	tmp := sl[0]
	tmp_idx := 0
	for i, v := range sl {
		if sl[i] > tmp {
			tmp = v
			tmp_idx = i
		}
	}
	fmt.Println("Q2==The max number:", tmp)
	fmt.Println("Q3==The origen slice:", sl)

	copy(sl[tmp_idx:], sl[tmp_idx+1:])
	sl[len(sl)-1] = tmp
	fmt.Println("Q3==The max follows the others:", sl)

	tmp = sl[0]
	tmp_idx = 0
	for i := 1; i < len(sl)-2; i++ {
		if sl[i] > tmp {
			tmp = sl[i]
		}
	}
	//
	copy(sl[tmp_idx:], sl[tmp_idx+1:])
	sl[len(sl)-2] = tmp
	fmt.Println("Q3==The second max follows the lessers:", sl)

	//
}
