package main

import "fmt"

func main()  {
	for i:=1;i<=9;i++{
		for k:=1;k<=i;k++{
			fmt.Printf("%d * %d = %d \t",k,i,i*k)
		}
		fmt.Println()
	}
}