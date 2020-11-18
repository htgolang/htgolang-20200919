package utils

import (
	"fmt"
	"testing"
)

func TestStrConversionTime_test(t *testing.T)  {
	data := StrConversionTime("1994-04-06 18:00:00")
	str := TimeConversionTimestamp(data)
	fmt.Printf("%T %v\n",str,str)
}