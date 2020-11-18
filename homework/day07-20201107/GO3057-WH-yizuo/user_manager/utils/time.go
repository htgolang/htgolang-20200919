package utils

import (
	"time"
)


// 将传递的字符串作为时间格式返回
func StrConversionTime(str string) (ret time.Time) {
	ret, _ = time.Parse("2006-01-02 15:04:05", str)
    return
}

// 将传递的时间格式转换为字符串
func TimeConversionTimestamp(ret time.Time) (str string) {
	str = ret.Format("2006-01-02 15:04:05")
	return
}