package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	txt := "i am kk"
	md5Value := fmt.Sprintf("%X", md5.Sum([]byte(txt)))
	fmt.Println(md5Value)
	sha1Value := fmt.Sprintf("%X", sha1.Sum([]byte(txt)))
	fmt.Println(sha1Value)
	sha256Value := fmt.Sprintf("%X", sha256.Sum256([]byte(txt)))
	fmt.Println(sha256Value)
	sha512Value := fmt.Sprintf("%X", sha512.Sum512([]byte(txt)))
	fmt.Println(sha512Value)

	md5hasher := md5.New()
	md5hasher.Write([]byte("i am"))
	md5hasher.Write([]byte(" kk"))

	// sha1Hasher := sha1.New()
	// sha1Hasher.Write([]byte("xx"))
	// sha1Hasher.Sum((nil))
	fmt.Printf("%X", md5hasher.Sum(nil))

}
