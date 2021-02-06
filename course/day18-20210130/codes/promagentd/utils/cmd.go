package utils

import "os/exec"

func Run(cmd string) (string, error) {
	command := exec.Command("/bin/bash", "-c", cmd)
	output, err := command.CombinedOutput()
	return string(output), err
}
