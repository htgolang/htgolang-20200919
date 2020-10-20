# 改
else if v < maxNum {}  不需要处理的话，要尽量省略这些逻辑，更简洁


if nums[i] > nums[j] { // 前面的值大于后面的值 进行交换
	temp := nums[i]      // 将nums[i] 的值赋予给temp，此时temp 为nums[i] 的初始值比较为2个数中的较大的值
	nums[i] = nums[j] // 将nums[j]的值 赋予给nums[i],此时nums[i] 为2个数中的较小值
	nums[j] = temp  // 将 temp(较大的值) 赋予给nums[j]，交换完成
}

数据交换，golang里有一个简洁写法
if nums[i] > nums[j] {
    nums[i],nums[j] = nums[j],nums[i]
}