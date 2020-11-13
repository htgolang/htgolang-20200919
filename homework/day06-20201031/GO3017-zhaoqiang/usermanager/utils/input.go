package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strings"

	"github.com/howeyc/gopass"
)

func Md5Convert(text []byte) string {
	m := md5.Sum(text)
	return fmt.Sprintf("%x", m)
}

func GetPasswd(prompt string) (string, error) {
	fmt.Printf("%s", prompt)
	bpass, err := gopass.GetPasswd()
	return string(bpass), err
}

func GetInput(prompt string) string {
	fmt.Printf("%s", prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
