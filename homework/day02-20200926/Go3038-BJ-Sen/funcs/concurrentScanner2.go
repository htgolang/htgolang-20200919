package funcs

import (
	"bufio"
	"fmt"
	"os"
)

func ConcurrentScannerFile() {
	println("open file")

	//p := filepath.Dir("day02/I have a dream")
	//println(p)
	file, err := os.Open("day02/I have a dream")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println("read lines:")
	for _, line := range lines {
		fmt.Println(line)
	}
	println("over")
}
