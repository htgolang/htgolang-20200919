package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// gen a id by UnixNano() who's type is int64
func GenId() (res int64) {
	result := time.Now().UnixNano()
	return result
}

// Message print debug info
func Message(s string) {
	fmt.Println(s)
}

// to verify if a string contains only digits
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

// read content from standard input
func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	return line
}

// gen [16]uint8 password
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
func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
