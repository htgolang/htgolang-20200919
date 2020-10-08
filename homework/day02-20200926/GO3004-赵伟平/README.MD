1. 统计"我有一个梦想"英文文章中每个英文字母(不区分大小写字母)(不统计标点符号)出现的次数
    map
    for
    关系运算
    strings

**代码：**
[wordFrequencyStatistics.go](./01wordFrequencyStatistics.go)

**结果：**
```
# go run 01wordFrequencyStatistics.go 
t: 671
n: 485
w: 165
u: 176
g: 175
z: 6
a: 567
y: 128
d: 270
r: 413
e: 889
b: 116
x: 5
q: 7
h: 384
j: 21
c: 182
m: 188
f: 224
o: 607
l: 342
s: 429
v: 81
k: 51
i: 565
p: 96
```

2. int切片 []int{108, 107, 105, 109, 103, 102}
    找出最大的数字

**代码：**
[02maxNum.go](./02maxNum.go)

**结果：**
```
# go run 02maxNum.go  
109
```

3. 针对问题2的切片, 将最大的数字移动到切片的最后一位
    原来的数字都在移动后的切片中都存在, 只是最大的数字再最后一位

**代码：**
[03mvMaxNum.go](./03mvMaxNum.go)

**结果：**
```
# go run 03mvMaxNum.go
[108 107 105 103 102 109]
```

4. 针对问题3, 将第二大的移动到切片的最后第二位

**代码：**
[04mvMaxSecondNum.go](./04mvMaxSecondNum.go)

**结果：**
```
# go run 04mvMaxSecondNum.go
[107 105 103 102 108 109]
```

5. 使用冒泡排序算法对问题2切片进行从小到大排序

提醒:
    int a, b = 1, 4
    交换方式一:
    int tmp = a
    a=b
    b = tmp
    交换方式二:
    a, b = b, a

**代码：**
[05BubbleSort.go](./05BubbleSort.go)

**结果：**
```
# go run 05BubbleSort.go
[102 103 105 107 108 109]
```

6. 挑战:
    a. 对问题2中的切片使用插入排序 实现从小到大排序
    b. 生成n个随机数, 在已排序的切片中(插入排序或者冒泡排序)查找其索引并打印, 为找到打印-1


