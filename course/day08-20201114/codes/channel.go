package main

import (
	"fmt"
	"time"
)

type Element struct {
}

func main() {
	// chan
	// channel

	// 管道中放什么类型需要提前执行
	// 声明int类型的管道channel
	// var channel chan int = make(chan int)
	channel := make(chan int) // 无缓冲区管道

	// 初始化&赋值
	// make()
	// channel = make(chan int)

	// 操作
	// 读，写
	go func() {
		fmt.Println("go start")
		channel <- 1 // 将1写入管道
		fmt.Println("go 1 end", time.Now())
		channel <- 2 // 将1写入管道
		fmt.Println("go end", time.Now())
	}()
	fmt.Println("channel begin")

	// <-channel // 读管道
	// num := <-channel
	// func test()  int {return 1}
	// test()
	// num := test()

	num, ok := <-channel         // 如果未读取到数据会进行阻塞
	fmt.Println("channel after") // go start 之后
	fmt.Println(num, ok)
	time.Sleep(3 * time.Second)
	<-channel
	time.Sleep(2 * time.Second)

	// 关闭管道
	// close(channel)

	// go func() {
	// fmt.Println(<-channel)
	// }()
	// 不能写
	// channel <- 1

	// go func() {
	// 	fmt.Println("close befoer")
	// 	time.Sleep(3 * time.Second)
	// 	close(channel)
	// }()

	// v, ok := <-channel
	// fmt.Println(v, ok)

	// 遍历
	go func() {
		channel <- 1
		channel <- 2
		channel <- 3
		close(channel)
	}()

	for num := range channel {
		fmt.Println(num)
	}
}
