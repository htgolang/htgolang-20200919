package main

import "fmt"

type EmailSender struct {
}

func (s *EmailSender) Send(msg string) error {
	fmt.Println("msg:", msg)
	return nil
}

func main() {

	var sender interface {
		Send(msg string) error
	}

	sender = new(EmailSender)
	sender.Send("xxxx")

}
