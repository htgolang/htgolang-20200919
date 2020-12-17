package main

import (
	"fileobj/utils"
	"fmt"
	"os"
)

func main() {
	ok := utils.CheckFileExist(os.Args[1])
	if !ok {
		fmt.Printf("%s文件找不到\n", os.Args[1])
	}
	filePerm := utils.GetFilePerm(os.Args[1])
	srcObj, _ := os.Open(os.Args[1])

	defer srcObj.Close()
	//做目标文件的判断
	ok = utils.CheckFileExist(os.Args[2])

	//说明目标文件不存在
	if !ok {
		fmt.Printf("%s不存在\n", os.Args[2])
		utils.HandleWriteFile(srcObj, os.Args[2], os.O_CREATE|os.O_WRONLY, filePerm)
	} else {
		//说明目标文件存在，询问是否覆盖，如果覆盖以os.Truncate形式和写文件形式打开
		var input string
		fmt.Printf("%s已经存在，是否覆盖? \n", os.Args[2])
		for {

			fmt.Scanln(&input)
			switch input {
			case "y":
				utils.HandleWriteFile(srcObj, os.Args[2], os.O_TRUNC|os.O_WRONLY, filePerm)
				os.Exit(0)
			case "n":
				fmt.Println("byebye")
				os.Exit(0)
			default:
				fmt.Println("请输入y/n")
			}
		}
	}

}
