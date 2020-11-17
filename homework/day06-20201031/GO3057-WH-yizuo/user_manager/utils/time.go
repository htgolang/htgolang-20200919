package utils

import (
	"time"
)


// 将传递的字符串作为时间格式返回
func TimeConversion(str string) (ret time.Time) {
	ret, _ = time.Parse("2006-01-02 15:04", str)
    return
}