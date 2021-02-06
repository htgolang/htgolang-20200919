package errors

import (
	"github.com/astaxie/beego/validation"
)

type Errors map[string][]string

func NewErrors() Errors {
	return make(Errors)
}

func (e Errors) AddError(key, err string) {
	e[key] = append(e[key], err)
}

func (e Errors) Clear() {
	keys := make([]string, 0, len(e))
	for k := range e {
		keys = append(keys, k)
	}
	for _, key := range keys {
		delete(e, key)
	}
}

func (e Errors) AddErrorMap(m map[string][]*validation.Error) {
	for _, ev := range m {
		for _, v := range ev {
			e.AddError("default", v.Error())
		}
	}
}
