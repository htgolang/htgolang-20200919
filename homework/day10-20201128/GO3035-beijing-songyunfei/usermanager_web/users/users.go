package users

import (
	"os"
)

var Savepath = ""
var Filefd *os.File
var QueueLen int
var queue [][]Userinfo
