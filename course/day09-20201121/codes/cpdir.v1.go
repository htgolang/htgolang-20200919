package main

import "sync"

func copyFile(src, dst string, wg *sync.WaitGroup) {
	wg.Done()
}

func listDir(dir string) []string {
	return nil
}

func main() {
	srcDir, dstDir := '', ''
	// 检查
	names := listDir(srcDir)
	wg := &sync.WaitGroup{}
	for name := range names {
		srcFile := ''
		dstFile := ''
		wg.Add(1)
		go copyFile(srcFile, dstFile, wg)
	}

	wg.Wait()
}
