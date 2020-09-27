package main

import (
	"math/rand"
	"time"
)

//RandomMap ...
func RandomMap(n int) (mp map[int]int) {
	rand.Seed(time.Now().Unix())
	mp = make(map[int]int)
	for i := 0; i < n; i++ {
		mp[rand.Intn(200)] = -1
	}
	return mp
}
