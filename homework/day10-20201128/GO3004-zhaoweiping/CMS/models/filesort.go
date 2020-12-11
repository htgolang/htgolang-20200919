package models

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"time"
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

func SortFile(path, name string) (files ByModTime) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fis, err := f.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
	files = make(ByModTime, len(fis)+10)
	j := 0
	for i, v := range fis {
		if strings.Contains(fis[i].Name(), name) {
			files[j] = v
			j++
		}
	}
	files = files[:j]
	// for i := 0; i < len(fis); i++ {
	// 	if !strings.Contains(fis[i].Name(), name) {
	// 		fis = append(fis[:i], fis[i+1:]...)
	// 		i--
	// 	}
	// }

	sort.Sort(ByModTime(files))
	// for _, fi := range files {
	// 	fmt.Println(fi.Name())
	// }
	return
}

func DealWithFiles(path, name string) (filename string) {
	timestamp := time.Now().Format("20060102.150405")
	filename = path + name + "." + timestamp

	files := SortFile(path, name)
	CopyFile(path+files[len(files)-1].Name(), filename)
	if len(files) > 5 {
		for k, _ := range files[:len(files)-5] {
			err := os.Remove(path + files[k].Name())
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if len(files) == 0 {
		f, err := os.Create(filename)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	// fmt.Println(filename)
	return filename
}

func CopyFile(file1, file2 string) {
	// 打开原始文件
	originalFile, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	_, err = io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Copied %d bytes.", bytesWritten)
	// 将文件内容flush到硬盘中
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
