package main

import (
	"fmt"
	process "vinash/process"
)

func main() {
	processes, err := process.GetProcesses()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, process := range processes {
		fmt.Println(process.Pid, process.Name)
	}
}
