package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ScanInt() (int, error) {
	// 读取一行 进行转换
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strconv.Atoi(scanner.Text())
	}
	return 0, scanner.Err()
}

func main() {
	// os.Stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	num, err := ScanInt()
	fmt.Println(num, err)

	num, err = ScanInt()
	fmt.Println(num, err)

	num, err = ScanInt()
	fmt.Println(num, err)

	num, err = ScanInt()
	fmt.Println(num, err)

	num, err = ScanInt()
	fmt.Println(num, err)
}
