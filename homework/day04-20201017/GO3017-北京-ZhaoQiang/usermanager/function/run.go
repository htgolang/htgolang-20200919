package function

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//OrderInput read input for user stdin, return user is input
//获取用户命令行输入
func OrderInput() string {
	scanner := bufio.NewScanner(os.Stdin)

START:
	fmt.Printf("[userManager] $ ")
	if scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			goto START
		}
		return scanner.Text()
	}
	return ""
}

//Run 运行命令行的命令
func Run(orderLine string) {
	order := strings.ToLower(orderLine)
	switch {
	case order == "add":
		AddRun()
	case order == "del":
		Del()
	case order == "modify":
		Modify()
	case order == "query":
		Query()
	case order == "h" || order == "help":
		View()
	case order == "printall":
		Printall()
	case order == "exit" || order == "quit" || order == "q":
		os.Exit(0)
	default:
		View()
		fmt.Println("输入提示的字符串")
	}

}
