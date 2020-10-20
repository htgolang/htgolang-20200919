package main

import (
	"testmultpkg/pkga"
	"testmultpkg/pkga/pkg"
	"testmultpkg/pkgb"
)

func main() {
	pkga.Test()
	pkgb.Test()
	pkg.Test()
}
