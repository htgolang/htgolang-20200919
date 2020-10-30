package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day05-20201024/Go3028-Beijing-lisuo/cmd/define"
)

// DateCheck make sure the input date is formatted
func DateCheck(d string) error {
	_, err := time.Parse("2006.01.02", d)
	return err
}

// GenID gen a id by UnixNano() who's type is int64
func GenID() (res int64) {
	result := time.Now().UnixNano()
	return result
}

// JustDigits to verify if a string contains only digits
func JustDigits(s string) bool {
	var a bool = true
	for _, c := range s {
		if c < '0' || c > '9' {
			a = false
			break
		}
	}
	return a
}

// Read read content from standard input
func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	return line
}

// GetField prompt input the User's field
func GetField(f string) string {
	for _, field := range define.UserField {
		if field == f {
			fmt.Printf("Please input %v: ", f)
			input := Read()
			return input
		}
	}
	return f
}

// GenPasswd gen [16]uint8 password
func GenPasswd() [16]uint8 {
	d := []byte(Read())
	return md5.Sum(d)
}

// clear the console
var clear map[string]func() //create a map for storing clear funcs
func init() {
	clear = make(map[string]func()) //Initialize it
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
}

// ClearScreen swip the chars on terminal
func ClearScreen() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// Message print debug info
func Message(v string) {
	fmt.Println(v)
}
