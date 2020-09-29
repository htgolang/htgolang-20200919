
package main

import (
"fmt"
)

func main () {
	sl := []int{108, 107, 105, 109, 103, 102}
	for i:=0;i<len(sl);i++{
	fmt.Println("++++",sl[i])
		for j:=1;j<len(sl);j++{
		fmt.Println("----",sl[j])
			if sl[i] > sl[j]{
				sl[i],sl[j]=sl[j],sl[i]
			}else{
				break
			}
		}
		fmt.Println(sl)
	}
	fmt.Println(sl)
}
