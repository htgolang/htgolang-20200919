package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

//CallClear 调用操作系统的clear清屏
func CallClear() {
	clear := make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

//Input 传入参数prompt--提示请输入什么,返回用户的输入的字符串
func Input(prompt string) string {
	fmt.Printf("%s ", prompt)
	text := "" //不论用户输入的是什么类型,强制转为字符串
	fmt.Scan(&text)
	return text
}

//GenID 在切片中增加一个ID,现有最大值+1
func GenID(s []int) int {
	if len(s) == 0 {
		rand.Seed(time.Now().Unix())
		return rand.Intn(100)
	}
	return s[len(s)-1] + 1
}

//PasswordEncrypt 传入明文密码,返回MD5密文
func PasswordEncrypt(pas string) string {
	return fmt.Sprintf("%X", md5.Sum([]byte(pas)))
}

//IsFileExist 判断文件路径是否存在,是否为目录
func IsFileExist(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		// fmt.Println(err)       //不存在,没权限等等
		return os.IsExist(err) //err不为空时才有必要使用
	}
	if p.Mode().IsDir() {
		fmt.Println("Attention!!", path, "Is DIR!!")
		return false
	}
	return true
}
