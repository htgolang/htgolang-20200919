package pkg

import "fmt"

type Sender interface {
	send(string) error
}

type EmailSender struct {
}

type Connection interface {
	Sender
}

func (s *EmailSender) send(msg string) error {
	fmt.Println("msg", msg)
	return nil
}

func Test() {
	var sender Sender = new(EmailSender)
	sender.send("xxx")
}
