package utils

import (
	"fmt"
	"os"
	"sort"
)

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().Before(fis[j].ModTime())
}

// 根目录下的文件按时间大小排序，从远到近
func SortFile(path string) (fis ByModTime) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fis, err = f.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	sort.Sort(ByModTime(fis))
	// for _, fi := range files {
	// 	fmt.Println(fi.Name())
	// }
	return fis
}

// func main() {
// 	path := "/Users/aomine/Documents/projects/scripts/go/htgolang-20200919/homework/day10-20201128/GO3004-zhaoweiping/uploadFile/files"
// 	files := SortFile(path)
// 	for _, v := range files {
// 		fmt.Println(v)
// 	}
// }
