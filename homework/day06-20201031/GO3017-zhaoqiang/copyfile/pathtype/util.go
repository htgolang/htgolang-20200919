package pathtype

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func copydir(source, dest string) {
	fileinfos, err := ioutil.ReadDir(source)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, fileinfo := range fileinfos {
		if fileinfo.IsDir() {
			// 是目录
			os.Mkdir(filepath.Join(dest, fileinfo.Name()), 0777)
			copydir(filepath.Join(source, fileinfo.Name()), filepath.Join(dest, fileinfo.Name()))
		} else {
			//是文件
			// copyfile(filepath.Join(source, fileinfo.Name()), filepath.Join(dest, fileinfo.Name()))
			if !overWrite(filepath.Join(source, fileinfo.Name()), dest) {
				copyfile(filepath.Join(source, fileinfo.Name()), filepath.Join(dest, fileinfo.Name()))
			}
		}
	}
}

func copyfile(source, dest string) {
	sfile, err := os.Open(source)
	dfile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sfile.Close()
	defer dfile.Close()

	bytes := make([]byte, 1024*1024*50)
	_, cberr := io.CopyBuffer(dfile, sfile, bytes)
	if cberr != nil {
		fmt.Println(err)
		return
	}
}

func overWriteBase(source, dest string, cmdcompare string) bool {
	var readdirdest string
	var overweitePrompt string

	info, _ := os.Stat(dest)
	if info.IsDir() {
		readdirdest = dest
		overweitePrompt = filepath.Join(dest, filepath.Base(source)) //dest是目录
	} else { //路径存在是文件
		readdirdest = filepath.Dir(dest)
		overweitePrompt = dest //dest是文件，提示dest的路径
	}

	fileinfos, _ := ioutil.ReadDir(readdirdest)
	for _, fileinfo := range fileinfos {
		if filepath.Base(cmdcompare) == fileinfo.Name() {
			fmt.Printf("cpfile overwrite %s? ", overweitePrompt)
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				if strings.ToLower(strings.TrimSpace(scanner.Text())) == "y" {
					var nowdest string
					if info.IsDir() {
						nowdest = filepath.Join(dest, filepath.Base(source))
					} else {
						nowdest = dest
					}
					// fmt.Println("overwrite222", source, nowdest)
					copyfile(source, nowdest)
					return true
				} else {
					return true
				}
			}
		}
	}
	return false
}

func overWriteFileFile(source, dest string) bool {
	return overWriteBase(source, dest, dest)
}

func overWriteFileDir(sourcefilepath, destDir string) bool {
	return overWriteBase(sourcefilepath, destDir, sourcefilepath)
}

func overWrite(source, dest string) bool {
	return overWriteBase(source, dest, source)
}
