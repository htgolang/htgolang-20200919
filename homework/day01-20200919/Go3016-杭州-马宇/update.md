# 改
\t的方式对齐很不错，

也可以尝试使用fmt.printf的占位试试看
fmt.Printf("%d*%d=%-2d ", m, i, m*i)

rand.Intn缺少随机数种子，自己尝试测试一下，看看如果没有随机数种子，是不是每次rand.Intn的时候，都是同一个数字，不会变