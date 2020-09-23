// 实现 99 乘法表
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			// fmt.Printf("%d x %d = %2d \n", j, i, i*j)
			if j == i {
				fmt.Printf("%d x %d = %2d \n", j, i, i*j)
			} else {
				fmt.Printf("%d x %d = %2d ", j, i, i*j)
			}

		}
	}
}

/*
1 双循环实现
2 有结果反推出，每一行的 被乘数是不变的，故这个数是放在外循环，从而在输出的时候也就定了i是被乘数
3 内循环分析出  乘数都是小于等于被乘数，故 得出 j <= i 的控制条件
4 输出的时候 当两个相乘的数 都相同的时候，就换行，故 通过条件判断 j == i 就分为两种不同的情况（换行和不换行），结果中有多位数，故采用占位符
*/
