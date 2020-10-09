package funcs

import (
	"fmt"
	"io"
	"os"
)

const BufferSize = 100

func BlockRead() {
	file, err := os.Open("day02/I have a dream")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		bytesread, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println("bytes read:", bytesread)
		fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	}
}
