package userutils

import "fmt"

func Input(str string) string {
	var tmp string
	fmt.Printf("%s",str)
	_,_= fmt.Scanln(&tmp)
	return tmp
}