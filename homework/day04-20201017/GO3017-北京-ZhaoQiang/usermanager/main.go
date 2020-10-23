package main

import (
	auth "zhao/authentication"
	"zhao/function"
)

func main() {
	var PasswdCount int = 3
	if auth.AuthUserPW(PasswdCount) {
		function.View()

		for {
			commlineInput := function.OrderInput()
			function.Run(commlineInput)
		}
	}

}
