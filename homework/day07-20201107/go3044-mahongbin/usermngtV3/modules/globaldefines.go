package modules

import (
	"time"
)

//Users 定义用户结构
type Users struct {
	ID                        int
	Name, Password, Tel, Addr string
	Birthday                  time.Time
	Deleted, Ifadmin          bool
}

//SliceU 保存全部用户的切片
var SliceU []Users

//SliceUID 保存全部用户ID的切片
var SliceUID []int

//DbFileInitType gob/csv/json
var DbFileInitType string

//DbFilePath 文件名路径
var DbFilePath string
