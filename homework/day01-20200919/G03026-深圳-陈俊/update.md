# 改
猜数字缺少随机种子rand.Seed(time.Now().Unix())

乘法口诀输出没有对齐，尝试使用fmt.printf的占位。
fmt.Printf("%d*%d=%-2d ", m, i, m*i)