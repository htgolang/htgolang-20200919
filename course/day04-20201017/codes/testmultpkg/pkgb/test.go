package pkgb

import (
	"fmt"
	"testmultpkg/pkga"
)

func Test() {
	fmt.Println("pkgb Test")
	pkga.Test()
}
