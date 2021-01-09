package main

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

func main() {
	// email := "kk@test.com"
	email := "test.com"
	valid := validation.Validation{}
	valid.Email(email, "email.email.email").Message("必须是邮箱格式")
	valid.MaxSize(email, 5, "email.email.email").Message("最大长度为5")
	fmt.Println(valid.HasErrors())
	fmt.Println(valid.Errors)
	fmt.Println(valid.ErrorMap())

	// Form
}
