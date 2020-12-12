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

type IPAndStatusCode struct {
	ip   string
	code int
}

type IPAndStatusCodeElement struct {
	key   IPAndStatusCode
	value int
}

type StringElement struct {
	key   string
	value int
}

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

	/*
		1.1.1.1 200
		1.1.1.1 200
		1.1.1.1 201
		1.1.1.2 201

		每个IP 每个状态码
		1.1.1.1 200 2
		1.1.1.1 201 1
		1.1.1.2 201 1
	*/
	file, err := os.Open("access.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	ipAndStatusCodeStats := make(map[IPAndStatusCode]int)

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
			key := IPAndStatusCode{ip, statusCode}
			ipAndStatusCodeStats[key]++
		}

	}

	fmt.Println(ipAndStatusCodeStats)

	list := make([]IPAndStatusCodeElement, 0, len(ipAndStatusCodeStats))
	for k, v := range ipAndStatusCodeStats {
		list = append(list, IPAndStatusCodeElement{k, v})
	}

	// 冒泡
	sort.Slice(list, func(i, j int) bool {
		return list[i].value < list[j].value
	})

	fmt.Println(list)
}
