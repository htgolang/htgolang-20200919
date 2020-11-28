package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	//strings, bytes
	// 内存字符串的操作
	// 读，写
	// strings.Reader
	reader := strings.NewReader("123abcxyz")
	fmt.Println(reader.Len(), reader.Size())

	ctx := make([]byte, 5)
	n, err := reader.Read(ctx)
	fmt.Println(n, err, string(ctx[:n]))
	fmt.Println(reader.Len(), reader.Size())

	n, err = reader.Read(ctx)

	fmt.Println(n, err, string(ctx[:n]))

	reader.Seek(0, 0)
	fmt.Println(reader.Len(), reader.Size())

	n, err = reader.Read(ctx)
	fmt.Println(n, err, string(ctx[:n]))

	reader.Reset("xyz")
	reader.WriteTo(os.Stdout)

	fmt.Println()
	// strings.Builder
	var builder strings.Builder
	builder.Write([]byte("Abc123"))
	fmt.Println(builder.Len())
	builder.WriteString("xyz")
	fmt.Println(builder.String())
	fmt.Println(builder.Len())
	builder.Reset()
	fmt.Println(builder.String())
	fmt.Println(builder.Len())

	// bytes
	// Reader, Buffer
	byteReader := bytes.NewReader([]byte("abc123abc"))
	fmt.Println(byteReader.Len(), byteReader.Size())

	n, err = byteReader.Read(ctx)
	fmt.Println(err, string(ctx[:n]))
	fmt.Println(byteReader.Len(), byteReader.Size())

	// Buffer 读/写
	buffer := bytes.NewBufferString("abc")
	buffer.WriteString("123")

	n, err = buffer.Read(ctx)
	fmt.Println(err, string(ctx[:n]))

	buffer.WriteString("xyz")

	fmt.Println(buffer.String())
}
