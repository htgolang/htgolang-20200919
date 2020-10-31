package main

import (
	"bufio"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	reader.WriteTo(os.Stdout)

	// content := make([]byte, 3)
	// n, err := reader.Read(content)
	// fmt.Println(n, err, content[:n])
	// fmt.Println(reader.ReadByte())
	// fmt.Println(reader.ReadBytes('|'))
	// fmt.Println(reader.ReadBytes('|'))
	// fmt.Println(reader.ReadBytes('|'))
	// fmt.Println(reader.ReadLine())
	// fmt.Println(reader.ReadLine())
	// fmt.Println(reader.ReadLine())
	// fmt.Println(reader.ReadLine())
	// fmt.Println(reader.ReadLine())
	// fmt.Println(reader.ReadSlice('|'))
	// fmt.Println(reader.ReadSlice('|'))
	// fmt.Println(reader.ReadSlice('|'))
	// fmt.Println(reader.ReadString('|'))
	// fmt.Println(reader.ReadString('|'))
	// fmt.Println(reader.ReadString('|'))
}
