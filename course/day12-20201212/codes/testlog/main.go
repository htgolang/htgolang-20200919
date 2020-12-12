package main

import (
	"fmt"
	"log"
	"os"
)

// docker container logs containername
func main() {
	logFile, _ := os.Create("test.log")
	defer logFile.Close()
	// log包
	// 日志格式
	// 前缀
	log.SetPrefix("testlog")
	fmt.Println(log.Flags())
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.SetOutput(logFile)

	log.Print("print")
	log.Println("println")
	log.Printf("%s xxx", "kk")
	// log.Panic("panic")
	// Panicln
	// Panicf
	// log.Fatal("fatal")
	// Fatalln
	// Fatalf
	log.Println("print2")

	//
	logger1 := log.New(os.Stdout, "logger1:", log.Ldate|log.Lshortfile)
	logger2 := log.New(os.Stdout, "logger2:", log.Ldate|log.Ltime|log.Llongfile)
	logger1.Print("logger1")
	logger2.Print("logger2")

}
