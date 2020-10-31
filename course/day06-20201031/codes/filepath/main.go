package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	path, _ := filepath.Abs("main.go")
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Dir(path))
	fmt.Println(filepath.Split(path))
	fmt.Println(filepath.SplitList("a;b;c"))
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.HasPrefix(path, "d:\\a"))
	fmt.Println(filepath.IsAbs(path), filepath.IsAbs("a.go"))

	fmt.Println(filepath.Glob("dir/a.*"))
	fmt.Println(filepath.Glob("dir/a.*"))
	fmt.Println(filepath.Glob("dir/*/*.go"))

	filepath.Walk("dir", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, info.Name())
		return nil
	})

	fmt.Println(filepath.Join("d:\\a", "b/c/d", "a.txt"))
}
