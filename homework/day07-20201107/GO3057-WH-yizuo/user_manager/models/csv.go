package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"yizuo/utils"
)

// Read user data from a CSV file
func ReadUsersDataToCsv()  {
	file, _ := os.Open(UserDataFile)
	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		/*
		//添加用户数据，迁移至user.go文件中进行所有与数据库有关的操作
		var (
			// 根据输入值插入数据
			u = &User{
				Id:       cast.ToInt(line[0]),             // 用户id
				Name:     line[2],                           // 名称
				Password: line[3],                           // 密码
				Phone:    line[4],                           // 联系方式
				Address:  line[5],                           // 地址
				Birthday: utils.StrConversionTime(line[6]),  // 生日
			}
		)

		var users = Users{
			curId:		cast.ToInt(line[0]),    // ID
			Status:		cast.ToInt(line[1]),    // 用户状态。0为软删除，1为在使用
			UserData:	u,    // 传入数据
		}
		UserList = append(UserList,users)
		*/
		ReadUserList(line[0],line[1],line[0],line[2],line[3],line[4],line[5],line[6])
	}
}

// User data is written to a CSV file
func WritesUsersDataToCsv()  {
	//User status
	userStatus,err := os.Create(UserDataFile)
	if err != nil {
		return
	}
	writer1 := csv.NewWriter(userStatus)
	for _,v := range UserList {
		_ = writer1.Write(
			[]string{
				strconv.Itoa(v.curId),
				strconv.Itoa(v.Status),
				v.UserData.Name,
				v.UserData.Password,
				v.UserData.Phone,
				v.UserData.Address,
				utils.TimeConversionTimestamp(v.UserData.Birthday),
			})
	}
	writer1.Flush()
	return
}