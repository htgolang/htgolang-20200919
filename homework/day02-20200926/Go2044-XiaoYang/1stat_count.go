package main

import (
	"fmt"
	"strings"
)

func main() {

	/*

		1.统计"我有一个梦想"英文文章中每个英文字母(不区分大小写字母)(不统计标点符号)出现的次数 map for 关系运算 strings



	*/

	info :=
		`I have a dream,a song to sing我有一个梦想，bai去唱一首歌
	To help me cope with anything以帮助du我去应对一切
	If you see the wonder of a fairy tale如果你看到了童zhi话中的奇迹
	You can take the future even if you fail即使你失败dao了，你也可以掌握未来
	I believe in angels我相信有天使存在
	Something good in everything I see我看见了一切美好的东西
	I believe in angels我相信有天使存在
	When I know the time is right for me当我知道时间对我说到了正确的时候
	I'll cross the stream我将穿越溪流
	I have a dream我有一个梦想
	I have a dream,a fantasy我有一个梦想，一个幻想
	To help me through reality去帮助我熬过现实
	And my destination makes it worth the while我的目标让那梦想有了价值
	Pushing through the darkness still another mile度过黑暗还有一英里
	I believe in angels我相信有天使存在
	Something good in everything I see我看见了一切美好的东西
	I believe in angels我相信有天使存在
	When I know the time is right for me当我知道时间对我说到了正确的时候
	I'll cross the stream 我将穿越溪流
	I have a dream我有一个梦想
	I have a dream,a song to sing我有一个梦想，去唱一首歌
	To help me cope with anything以帮助我去应对一切
	If you see the wonder of a fairy tale如果你看到了童话中的奇迹
	You can take the future even if you fail即使你失败了，你也可以掌握未来
	I believe in angels我相信有天使存在
	Something good in everything I see我看见了一切美好的东西
	I believe in angels我相信有天使存在
	When I know the time is right for me当我知道时间对我说到了正确的时候
	I'll cross the stream我将穿越溪流
	I have a dream我有一个梦想
	I'll cross the stream我将穿越溪流
	I have a dream我有一个梦想
	`

	// 定义map
	var countstring map[string]int

	// 初始化
	countstring = map[string]int{}

	// countstring := make(map[string]int)

	for _, v := range info {

		s1 := string(v)
		// 字符串转小写
		s1 = strings.ToLower(s1)
		if s1 >= "a" && s1 <= "z" || s1 >= "A" && s1 <= "Z" {
			countstring[s1]++
		}
	}
	// fmt.Println(countstring)
	for char, count := range countstring {
		fmt.Printf("字符串: %s\t 出现次数:%d\n", char, count)
	}
}
