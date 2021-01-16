package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	// "ls -la /"
	// dir .
	// cmd := exec.Command("ls", "-la", "/")
	fmt.Println(exec.LookPath("ping"))
	cmd := exec.Command("ping", "www.xxx.com", "-c", "1")
	// output, err := cmd.Output()
	stdout, _ := cmd.StdoutPipe()

	cmd.Start()
	fmt.Println("started")

	io.Copy(os.Stdout, stdout)
	cmd.Wait()

	fmt.Println(cmd.ProcessState.ExitCode())
	os.Exit(100)

	// exex.Command(), Output()  cmd.ProcessState.ExitCode()
	// ->sh
}
