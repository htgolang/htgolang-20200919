package funcs

import (
	"fmt"
	"io/ioutil"
	"log"
)

/*
ASCII:
a-z: 97-122
A-Z: 65-90
*/

// CharNumsInDream get contents of I_have_a_dream.txt and
// calculate the alphabet letter count in the file
func CharNumsInDream() {
	var res = make(map[byte]int)
	var file = "./files/I_have_a_dream.txt"

	// read the whole file to data as byte
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// get a-z and A-Z, make all A-Z to a-z,
	// and make all others to NUL('\0')
	for i, asc := range data {
		if 64 < asc && asc < 91 {
			data[i] = asc + 32
		} else if asc < 97 || asc > 122 {
			data[i] = 0
		}
	}

	// accumulate the letter count
	for _, letter := range data {
		res[letter]++
	}

	for i, v := range res {
		if i != 0 {
			fmt.Printf("letter: %c, count: %v\n", i, v)
		}
	}

}
