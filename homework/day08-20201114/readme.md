1. cp 复制文件夹
    并发复制
    N个文件-> N个例程 waitgroup
    固定M个例程 做拷贝(worker) // waitgroup, channel

2. 遍历文件夹中文件, 过滤，统计go文件代码行数
    并发统计
    N个文件->N个例程
    固定M个例程 做计算(worker) // waitgroup, channel

3. 用户管理