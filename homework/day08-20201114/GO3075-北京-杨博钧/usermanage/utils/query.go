package utils

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

func Query() {
	sort.Sort(UsersList)
	ind := []int{}
	value := ""
	fmt.Print("请输入要查询的值:")
	fmt.Scan(&value)
	num, err := strconv.Atoi(value)
	if err == nil {
		for i, user := range UsersList {
			if num == user.Id || value == user.Tel {
				ind = append(ind, i + 1)
			}
		}
	}
	date, err := time.Parse("2006-01-02", value)
	if err == nil {
		for i, user := range UsersList {
			if value == user.Addr || value == user.Name ||
				date == user.Birthday {
				ind = append(ind, i + 1)
			}
		}
	} else {
		for i, user := range UsersList {
			if value == user.Addr || value == user.Name {
				ind = append(ind, i + 1)
			}
		}
	}
	if len(ind) > 0 {
		fmt.Println("匹配到了如下数据:")
		for _, i := range ind {
			if i >= 1 {
				fmt.Printf("%v\n", UsersList[i - 1])
			}
		}
	} else {
		fmt.Println("无匹配数据")
	}
}
