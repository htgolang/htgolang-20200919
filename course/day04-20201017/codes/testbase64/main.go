package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// 编码 []byte => string
	// 解码 string => []byte
	text := "i am kk编码"
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(text)))
	bytes, _ := base64.StdEncoding.DecodeString("aSBhbSBraw==")
	fmt.Println(string(bytes))
	fmt.Println(base64.RawStdEncoding.EncodeToString([]byte(text)))
	fmt.Println(base64.URLEncoding.EncodeToString([]byte(text)))
}
