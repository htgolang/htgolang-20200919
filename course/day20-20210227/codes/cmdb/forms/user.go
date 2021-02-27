package forms

import (
	"cmdb/services"
	"strings"

	"github.com/astaxie/beego/validation"
)

type AddUserForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
	Addr     string `form:"addr"`
	Sex      bool   `form:"sex"`
}

func (f *AddUserForm) Valid(v *validation.Validation) {
	f.Name = strings.TrimSpace(f.Name)
	v.Required(f.Name, "name.name.name").Message("用户名不能为空")
	f.Password = strings.TrimSpace(f.Password)

	v.MaxSize(f.Password, 32, "password.password.password").Message("密码长度必须在6~32位之间")
	v.MinSize(f.Password, 6, "password.password.password").Message("密码长度必须在6~32位之间")

	f.Addr = strings.TrimSpace(f.Addr)
	if f.Addr != "北京" && f.Addr != "西安" {
		v.SetError("addr", "地址必须为西安或北京")
	}

	if _, ok := v.ErrorsMap["name"]; !ok {
		if user := services.GetUserByName(f.Name); user != nil {
			v.SetError("name", "用户名已存在")
		}
	}

}
