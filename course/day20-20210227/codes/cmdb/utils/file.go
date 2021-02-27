package utils

import (
	"fmt"
	"os"
)

func Mkdir(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(path, os.ModePerm)
		}
		return fmt.Errorf("mkdir error")
	}
	return nil
}
