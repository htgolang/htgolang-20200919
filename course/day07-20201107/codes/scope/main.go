package main

import (
	"fmt"
	"scope/pkg"
)

type EmailSender struct {
}

func (s *EmailSender) send(msg string) error {
	fmt.Println("msg", msg)
	return nil
}

func main() {
	// var sender pkg.Sender = new(EmailSender)
	// fmt.Println(sender)
	pkg.Test()
}
