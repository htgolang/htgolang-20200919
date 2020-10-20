package funcs

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadDir() {
	filelist, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	println(filelist)

	for _, fileinfo := range filelist {
		if fileinfo.Mode().IsRegular() {
			bytes, err := ioutil.ReadFile(fileinfo.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Bytes read:", len(bytes))
			fmt.Println("String read:", string(bytes))
		}
	}
}
