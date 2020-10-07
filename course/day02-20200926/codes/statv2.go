package main

import "fmt"

func main() {
	// 投票结果
	names := []string{
		"小明",
		"小红",
		"小花",
		"小黑",
		"大黄",
		"小明",
		"大黄",
		"小黑",
		"大黄",
		"大黄",
	}
	/*
		小明 1+1
		小红 1
		小花 1
		小黑 1+1
		大黄 1+1+1+1
	*/
	stat := map[string]int{}
	for _, name := range names {
		stat[name]++
	}
	fmt.Println(stat)
}
