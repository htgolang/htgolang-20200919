package main

import "fmt"

func main() {
	left, right := "梨", "苹果"
	fmt.Println("left=", left, "right=", right)
	desktop := left // 梨放在桌子上
	left = right    // 将右手的苹果放在左手
	right = desktop // 右手从桌子上拿梨

	fmt.Println("left=", left, "right=", right)

	left, right = right, left
	fmt.Println("left=", left, "right=", right)
}
