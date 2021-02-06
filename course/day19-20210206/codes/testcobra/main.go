package main

import (
	"testcobra/cmd"
)

// all verbose ==> rootCmd
// web --port => webCmd
// api --host => apiCmd
func main() {
	cmd.Execute()
}
