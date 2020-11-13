package users

import (
	"os"
)

var Savepath = ""
var Filefd *os.File
var QueueLen = 3
var queue [][]Userinfo
