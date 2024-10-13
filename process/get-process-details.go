package process

import (
	"os/exec"
)

func GetDetails(pid string) (string, error) {
	cmd := exec.Command("ps", "-p", pid, "-o", "pid,comm,%mem,%cpu")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
