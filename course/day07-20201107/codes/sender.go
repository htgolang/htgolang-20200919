package main

import "fmt"

// 产生告警 => 发送给相关的人
// email, sms, weixin, dingding
// 发送的信息 => 发送的方式

type Sender interface {
	Send(string) error
}

type AllSender interface {
	Send(string) error
	SendAll([]string) error
}

type SmsSender struct {
}

func (s *SmsSender) Send(msg string) error {
	return nil
}

type EmailSender struct {
	Addr     string
	Port     string
	User     string
	Password string
	to       string
}

func (s *EmailSender) Send(msg string) error {
	fmt.Println("email sender:" + msg)
	return nil
}

func (s *EmailSender) SendAll(msg []string) error {
	fmt.Println("email sender all:", msg)
	return nil
}

func test(sender *DingDingSender) {
	sender.Send("yyy")
}

func test2(sender Sender) {
	sender.Send("yyyy")
}

type DingDingSender struct {
}

func (s *DingDingSender) Send(msg string) error {
	fmt.Println("dingding sender:" + msg)
	return nil
}

func (s *DingDingSender) SendAll(msg []string) error {
	fmt.Println("dingding sender all:", msg)
	return nil
}

func main() {
	// var sender *DingDingSender = new(DingDingSender)
	// sender.Send("xx")
	// test(sender)

	var sender Sender = new(EmailSender)
	var allsender AllSender = new(EmailSender)
	// sender := new(EmailSender)
	// sender 类型
	// fmt.Println(sender.Addr)
	// sender.SendAll([]string{""})
	allsender.Send("啦啦啦啦啦")
	allsender.SendAll([]string{"aaaaa"})

	// sender.Send("xxx")
	// test2(sender)
	// sender, allsender
	// allsender = sender
	sender = allsender
	sender.Send("xxxx")
	fmt.Println(allsender, sender)
}
