package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/htgolang/htgolang-20200919/tree/master/homework/day10-20201128/Go3028-Beijing-lisuo/user_manager/utils"
)

// Rotate do the rotate things
func Rotate() {
	fileNameHandle()
}

func fileNameHandle() {
	var timeStampToFileName = make(map[int64]string, 5)
	var timeStamp []int64
	var backups = 5

	fmt.Println("dbDir in fileNameHandle: ", filepath.Join(dbDir, SaveFlag))
	filesInfo, err := ioutil.ReadDir(filepath.Join(dbDir, SaveFlag))
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, file := range filesInfo {
		fmt.Println(file.Name())
		fileNameSlice := strings.Split(file.Name(), ".")
		timestamp, errParse := strconv.ParseInt(fileNameSlice[len(fileNameSlice)-1], 10, 64)
		if errParse != nil {
			//fmt.Println(errParse)
			continue
		}
		timeStamp = append(timeStamp, timestamp)
		timeStampToFileName[timestamp] = file.Name()
	}
	//utils.SuffleInt64Slice(&timeStamp)
	//fmt.Println(timeStamp)
	// desc sort timeStamp
	utils.SortInt64Slice(&timeStamp)
	//fmt.Println("after sort: ", timeStamp)
	for index, time := range timeStamp {
		if index > backups-1 {
			// get files to del
			path := filepath.Join(dbDir, SaveFlag)
			if err := os.Remove(filepath.Join(path, timeStampToFileName[time])); err != nil {
				log.Fatal(err)
			}
			//fmt.Println(timeStampToFileName[time])
		}
	}
}

func genFileNameSuffix() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02d_%02d-%02d-%02d.%d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.UnixNano())
}

func getFileNameList() {

}
