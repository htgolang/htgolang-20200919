package main

import "fmt"

type Sender interface {
	Send(string) error
}

type AllSender interface {
	Send(string) error
	SendAll([]string) error
}

type EmailSender struct {
	Addr string
	To   string
}

func (s *EmailSender) Send(msg string) error {
	fmt.Println("email sender:", msg)
	return nil
}

func (s *EmailSender) SendAll(msgs []string) error {
	fmt.Println("email sender:", msgs)
	return nil
}

func (s *EmailSender) Test() {
	fmt.Println("test")
}

type DingDingSender struct {
	To string
}

func (s *DingDingSender) Send(msg string) error {
	fmt.Println("dingding sender:" + msg)
	return nil
}

func test(sender Sender) {
	if obj, ok := sender.(*EmailSender); ok {
		fmt.Println("email sender", obj.Addr)
	} else if obj, ok := sender.(*DingDingSender); ok {
		fmt.Println("dingding sender", obj.To)
	} else {
		fmt.Println("error")
	}
}

func query(sender Sender) {
	// 类型查询
	// switch case + 接口变量.(type)

	switch obj := sender.(type) {
	case *EmailSender:
		fmt.Println("email sender", obj.Addr, obj.To)
	case *DingDingSender:
		fmt.Println("dingding sender", obj.To)
	default:
		fmt.Println("error")
	}
}

func main() {
	var sender Sender = &EmailSender{"xxx", "yyy"}

	// fmt.Println(sender.Addr)
	// fmt.Println(sender.To)
	// sender.Test()
	sender.Send("xxx")

	// 断言
	obj, ok := sender.(*EmailSender)
	fmt.Printf("%T, %#v, %#v\n", obj, obj, ok)
	if ok {
		fmt.Println(obj.Addr)
		fmt.Println(obj.To)
		obj.Test()

	}

	dingding, ok := sender.(*DingDingSender)
	fmt.Println(ok, dingding)

	test(sender)
	test(new(DingDingSender))

	query(sender)
	query(new(DingDingSender))

	var allsender AllSender = &EmailSender{"111", "222"}
	// allsender.Send()
	// allsender.SendAll()
	sender = allsender
	sender.Send("xxxx")

	allsender2, ok := sender.(AllSender)
	fmt.Println(allsender2, ok)
	allsender2.SendAll([]string{"xxxx", "xxxxxxxxxx"})

	esender, ok := sender.(*EmailSender)
	fmt.Println(esender, ok)
	fmt.Println(esender.Addr, esender.To)
	esender.Test()

	switch sender.(type) {
	case *EmailSender:
		fmt.Println("emailsender")
	case AllSender:
		fmt.Println("allsender")
	default:
		fmt.Println("default")
	}
	// fmt.Println(sender.(type))
}
