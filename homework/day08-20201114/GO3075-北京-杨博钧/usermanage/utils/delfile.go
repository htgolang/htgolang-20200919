package utils

import "os"

func DelFile() {
	fileInfos := GetSortedFileList()
	if len(fileInfos) > 5 {
		for _, file := range fileInfos[:len(fileInfos) - 5] {
			os.Remove("data/" + file.Name())
		}
	}
}
