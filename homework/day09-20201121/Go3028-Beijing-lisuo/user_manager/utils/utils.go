package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day09-20201121/Go3028-Beijing-lisuo/user_manager/define"
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
	line := strings.TrimSpace(scanner.Text())
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

// GetKeyByValue return the same value's different keys
func GetKeyByValue(m map[string]string, value string) []string {
	var keys []string
	for k, v := range m {
		if v == value {
			keys = append(keys, k)
		}
	}
	return keys
}

// ArrayToString convert a string array to string
func ArrayToString(s []string) string {
	// b is a type of []uint8
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
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

// SortInt64Slice sort a slice
func SortInt64Slice(s *[]int64) {
	for i := 0; i < len(*s); i++ {
		// if all the number shifted, sweep(nums, i) will return false
		if !sweep(s, i) {
			return
		}
	}
}

func sweep(nums *[]int64, passesDone int) bool {
	var firstIndex, secondIndex int = 0, 1
	doSwap := false
	// those already shifted is passesDone, minus them
	for secondIndex < len(*nums)-passesDone {
		var firstNum = (*nums)[firstIndex]
		var secondNum = (*nums)[secondIndex]
		// shift big one backward
		if firstNum < secondNum {
			(*nums)[firstIndex], (*nums)[secondIndex] = secondNum, firstNum
			doSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return doSwap
}

// SuffleInt64Slice suffles a slice
func SuffleInt64Slice(s *[]int64) {
	rand.Seed(time.Now().UnixNano())
	for i := len(*s) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}

}

// Message print debug info
func Message(v string) {
	fmt.Println(v)
}

// Quit cause program exit
func Quit() {
	os.Exit(define.UserQuit)
}
