package controller

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"zhao/models"
)

//View 命令视图
func View() {
	fmt.Println(`
	*********************************
	|	add	添加用户	|
	|	del	删除用户	|
	|	modify	修改用户	|
	|	query	查找用户	|
	|	print	打印用户	|	
	|	exit	退出		|
	|	help	帮助		|
	********************************* 
	`)
}

//CommandInput 获取命令行用户输入指令
func CommandInput() string {
	fmt.Printf("[usermanager-%s]$: ", models.LoginUser.Name)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			if strings.TrimSpace(scanner.Text()) == "" {
				fmt.Printf("[usermanager-%s]$: ", models.LoginUser.Name)
				continue
			} else {
				return strings.TrimSpace(scanner.Text())
			}
		}
	}
}
