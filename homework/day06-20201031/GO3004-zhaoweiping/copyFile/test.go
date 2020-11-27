package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ColorGroup struct {
	id     int
	name   string
	colors []string
}

func main() {
	group := ColorGroup{
		id:     1,
		name:   "Reds",
		colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	//生成json文件
	err = ioutil.WriteFile("test.json", b, os.ModeAppend)
	if err != nil {
		return
	}
	var data interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("data", data) //map[ID:1 Name:Reds Colors:[Crimson Red Ruby Maroon]]
	dataJson := data.(map[string]interface{})["Colors"]
	fmt.Println("dataJson", dataJson) //[Crimson Red Ruby Maroon]
	b11, err := json.Marshal(dataJson)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = ioutil.WriteFile("test11.json", b11, os.ModeAppend)
	if err != nil {
		return
	}
}
