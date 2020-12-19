package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "123@456"
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	fmt.Println(string(hashed), err)

	hashed1 := "$2a$16$2WnQ0ATTXRHhJ.VyovOL7.5n/WHgiFaU/tFAd7bvUFpTMkcta8crS"
	hashed2 := "$2a$16$qlRQEeF8uzZjnJrEJh..sumZ9J.D/femgOub7RxKYNm56R09qB9TC"

	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed1), []byte(password)))
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed2), []byte(password)))
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)))

	password = "123"
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed1), []byte(password)))
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed2), []byte(password)))
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)))
}
