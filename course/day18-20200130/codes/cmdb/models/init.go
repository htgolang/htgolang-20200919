package models

import "github.com/astaxie/beego/orm"

func init() {
	orm.RegisterModel(new(User), new(Alert), new(Agent))
}
