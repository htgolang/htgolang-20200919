package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// ip, url, status
func parse(line string) (string, string, int, error) {
	elements := strings.Fields(line)
	if len(elements) < 8 {
		return "", "", 0, fmt.Errorf("format error")
	}
	statusCode, err := strconv.Atoi(elements[8])
	if err != nil {
		return "", "", 0, err
	}
	return elements[0], elements[6], statusCode, nil
}

func main() {
	file, err := os.Open("access.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	ipStats := make(map[string]int)
	// urlStats := make(map[string]int)
	statusCodeStats := make(map[int]int)

	for {
		ctx, _, err := reader.ReadLine()
		// fmt.Println(err, string(ctx))
		// A
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		// B
		if ip, _, statusCode, err := parse(string(ctx)); err == nil {
			ipStats[ip]++
			// urlStats[url]++
			statusCodeStats[statusCode]++
		}

	}

	// fmt.Println(ipStats)
	// fmt.Println(urlStats)
	fmt.Println(statusCodeStats)

	ipStatsSlice := make([][]interface{}, 0, len(ipStats))
	for k, v := range ipStats {
		ipStatsSlice = append(ipStatsSlice, []interface{}{k, v})
	}

	// 冒泡
	sort.Slice(ipStatsSlice, func(i, j int) bool {
		count1 := ipStatsSlice[i][1].(int)
		count2 := ipStatsSlice[j][1].(int)
		return count1 < count2
	})

	topn := 10
	for i := len(ipStatsSlice) - 1; i >= 0 && i >= len(ipStatsSlice)-topn; i-- {
		fmt.Println(ipStatsSlice[i])
	}
}
