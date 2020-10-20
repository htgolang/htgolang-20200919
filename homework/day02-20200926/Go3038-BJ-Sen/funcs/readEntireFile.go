package funcs

import (
	"fmt"
	"os"
)

func ReadEntireFile() {
	file, err := os.Open("I have a dream")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileInfo.Size()
	buffer := make([]byte, filesize)

	bytesReader, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("bytes read:", bytesReader)
	fmt.Println("bytestream to string:", string(buffer))
}
