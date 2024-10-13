package process

import (
	"os/exec"
	"strconv"
	"strings"
)

func GetDetails(pid string) (ProcessDetails, error) {
	cmd := exec.Command("ps", "-p", pid, "-o", "pid,comm,%mem,%cpu,uid,gid,state")
	out, err := cmd.Output()
	if err != nil {
		return ProcessDetails{}, err
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		return ProcessDetails{}, nil
	}

	fields := strings.Fields(lines[1])
	if len(fields) < 7 {
		return ProcessDetails{}, nil
	}

	memory, _ := strconv.ParseFloat(fields[2], 64)
	cpu, _ := strconv.ParseFloat(fields[3], 64)

	processDetails := ProcessDetails{
		PID:     fields[0],
		Command: fields[1],
		Memory:  memory,
		CPU:     cpu,
		UID:     fields[4],
		GID:     fields[5],
		State:   fields[6],
	}

	return processDetails, nil
}
