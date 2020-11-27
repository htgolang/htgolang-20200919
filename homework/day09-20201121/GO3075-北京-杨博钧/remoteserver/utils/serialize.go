package utils

import (
	"encoding/json"
	"fmt"
	"remoteserver/model"
)
// json转化请求参数
func SerializeParams(params *model.Params) []byte {
	request, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	request = append(request, '\n')
	return request
}

// 解析json请求
func DeserializeParams(request []byte) *model.Params {
	var params model.Params
	err := json.Unmarshal(request, &params)
	if err != nil {
		fmt.Println(err)
	}
	return &params
}

// json格式化返回值
func SerializeReturns(returnvalue model.ReturnValue) []byte {
	request, err := json.Marshal(returnvalue)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	request = append(request, '\n')
	return request
}

// 解析json返回
func DeserializeReturns(request []byte) *model.ReturnValue {
	var params model.ReturnValue
	err := json.Unmarshal(request, &params)
	if err != nil {
		fmt.Println(err)
	}
	return &params
}