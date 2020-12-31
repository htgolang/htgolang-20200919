package services

import (
	"bufio"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"webtool/models"
)

func Upload(fh *multipart.FileHeader, uploadFile string) error {
	fmt.Printf("uploadFile: %#v\n", uploadFile)
	newFile, err := os.Create(uploadFile)
	fmt.Printf("created: %#v\n", newFile)
	if err != nil {
		fmt.Println("create error...")
		fmt.Println(err)
		return err
	}
	file, err := fh.Open()
	fmt.Printf("fh.Open: %#v", file)
	if err != nil {
		fmt.Println("fh.Open error...")
		fmt.Println(err)
		return err
	}
	defer file.Close()
	_, errc := io.Copy(newFile, file)
	if errc != nil {
		fmt.Println("copy error....")
		fmt.Println(err)
		return err
	}
	return nil
}

func CalcResult() ([]models.Result, error) {
	// calculate the result
	var err error
	models.IpList, models.StatusList, models.URLList, err = GetList(models.UploadFile)
	if err != nil {
		return nil, err
	}
	ipNum := GetNum(models.IpList)
	models.IpRankList = GetRank(ipNum)
	statusNum := GetNum(models.StatusList)
	statusRankList := GetRank(statusNum)
	// get methodList from URLList(request Line)
	for _, url := range models.URLList {
		reqLine := strings.Split(url, " ")
		if strings.Contains(reqLine[0], `\x`) || strings.Contains(reqLine[0], `tmp`) ||
			strings.Contains(reqLine[0], `WIN`) {
			continue
		}
		models.MethodList = append(models.MethodList, reqLine[0])
	}
	methodNum := GetNum(models.MethodList)
	methodRankList := GetRank(methodNum)
	resultList := NewResult()
	// put ip to rank
	for i := 1; i <= models.RankLen; i++ {
		for k, v := range models.IpRankList {
			resultList[models.RankLen-i].IP = v
			resultList[models.RankLen-i].IPNum = k
		}
	}
	// put status code to rank
	for i := 1; i <= models.RankLen; i++ {
		for k, v := range statusRankList {
			s, _ := strconv.Atoi(v)
			resultList[models.RankLen-i].Status = s
			resultList[models.RankLen-i].StatusNum = k
		}
	}
	// put method to rank
	for i := 1; i <= models.RankLen; i++ {
		for k, v := range methodRankList {
			resultList[models.RankLen-i].Method = v
			resultList[models.RankLen-i].MethodNum = k
		}
	}
	return resultList, nil
}

func NewResult() []models.Result {
	var rList = []models.Result{}
	for i := 1; i <= models.RankLen; i++ {
		r := models.Result{Rank: i}
		rList = append(rList, r)
	}
	return rList
}

func GetList(file string) (ipList, statusLit, URLList []string, err error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, nil, nil, err
	}
	//io.Copy(os.Stdout, f)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ipLine := strings.Split(line, " ")

		ipList = append(ipList, ipLine[0])
		models.StatusList = append(models.StatusList, ipLine[8])

		URLLine := strings.Split(line, "\"")
		URLList = append(URLList, URLLine[1])
	}
	return models.IpList, models.StatusList, models.URLList, nil
}

func GetNum(itermList []string) map[string]int {
	var itermNum = make(map[string]int)
	for _, iterm := range itermList {
		if _, ok := itermNum[iterm]; ok {
			itermNum[iterm]++
		} else {
			itermNum[iterm] = 1
		}
	}
	return itermNum
}

func GetRank(itermNum map[string]int) map[int]string {
	var rankList = make(map[int]string)
	for i := 0; i < models.RankLen; i++ {
		var max int
		var keyMax string
		// get max occured ip and the occurence
		for k, v := range itermNum {
			if v > max {
				max = v
				keyMax = k
			}
		}
		// put the max to rankList
		rankList[max] = keyMax
		// make the max to zero
		itermNum[keyMax] = 0
	}
	return rankList
}
