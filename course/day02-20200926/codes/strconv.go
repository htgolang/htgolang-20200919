package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串 <=> int
	i, err := strconv.Atoi("abc")
	fmt.Println(i, err)
	i, err = strconv.Atoi("-15")
	fmt.Println(i, err)
	fmt.Printf("%q\n", strconv.Itoa(45))
	// 字符串 <=> bool
	fmt.Printf("%q\n", strconv.FormatBool(true))
	fmt.Printf("%q\n", strconv.FormatBool(false))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("0"))
	fmt.Println(strconv.ParseBool("xxx"))
	// 字符串 <=> float
	fmt.Println(strconv.FormatFloat(1.68, 'E', -1, 32))
	fmt.Println(strconv.FormatFloat(1.68, 'f', -1, 32))
	fmt.Println(strconv.ParseFloat("1.34", 32))
	fmt.Println(strconv.ParseFloat("1.34", 64))

}
