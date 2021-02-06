package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	if ctx, err := ioutil.ReadFile(path); err != nil {
		return ""
	} else {
		return string(ctx)
	}
}

func WriteFile(path string, txt string) {
	ioutil.WriteFile(path, []byte(txt), os.ModePerm)
}
