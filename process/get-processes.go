package process

import (
	"os/exec"
	"strings"
)
func GetProcesses() ([]Process, error) {
	var outputProcesses []Process
	cmd := exec.Command("ps", "-e", "-o", "pid,comm,%mem", "--sort=-%mem")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	processes := strings.Split(string(out), "\n")

	for _, process := range processes[1 : len(processes)-1] {
		fields := strings.Fields(process)
		if len(fields) >= 2 {
			pid := fields[0]
			name := fields[1]
			outputProcesses = append(outputProcesses, Process{pid, name})
		}
	}
	return outputProcesses, nil
}
