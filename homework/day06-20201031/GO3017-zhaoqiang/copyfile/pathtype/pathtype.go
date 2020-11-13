package pathtype

import (
	"fmt"
	"os"
	"path/filepath"
)

func Run(source, dest string) {
	fileinfo, err := os.Stat(source)
	if err != nil {
		fmt.Println("pathtype/Run", err)
		return
	}
	if fileinfo.IsDir() {
		sourceIsDir(source, dest)
	} else {
		sourceIsFile(source, dest)

	}
}

func sourceIsDir(source, dest string) {
	// dest 目录下是否有和source目录重复的文件名字
	file, err := os.Open(dest)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("sourceIsDir", err)
		return
	}
	if fileinfo.IsDir() {
		newdest := filepath.Join(dest, filepath.Base(source))
		os.Mkdir(newdest, 0777) //创建一个以source目录命名的目录， 如果存在

		fileinfoss, _ := file.Readdir(-1)
		for _, finfo := range fileinfoss {
			if !finfo.IsDir() {
				if finfo.Name() == filepath.Base(newdest) {
					fmt.Println("destination path has same name file")
					return
				}
			}
		}
		copydir(source, newdest)
	} else {
		//source 是目录， dest是文件
		fmt.Println("destination path is file")
		return
	}
}

func sourceIsFile(source, dest string) {
	file, err := os.Open(dest)
	if err != nil {
		if os.IsNotExist(err) {
			copyfile(source, dest)
			return
		}
		fmt.Println("sourceIsfil", err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println("sourceIsFile", err)
		return
	}
	if fileinfo.IsDir() {
		//dest是目录
		if !overWriteFileDir(source, dest) { //dest目录下有相同名字的文件
			copyfile(source, filepath.Join(dest, filepath.Base(source)))
		}
	} else { //文件存在，有相同的名字  dest下有相同名字的文件
		if !overWriteFileFile(source, dest) {
			copyfile(source, dest)
		}

	}
}
