package main

import "sync"

func copyFile(src, dst string, wg *sync.WaitGroup) {
	wg.Done()
}

func listDir(srcDir string, dstDir string, callback func(srcFile, dstFile string)) []string {
	//
	// 遍历文件
	name => srcFile,
	name => dstFile
	callback(srcFile, dstFile)
	return nil
}

func main() {
	srcDir, dstDir := '', ''

	wg := &sync.WaitGroup{}
	// 检查
	names := listDir(srcDir, func(srcFile, dstFile string) {
		wg.Add(1)
		go copyFile(srcFile, dstFile)
	})

	wg.Wait()
}
