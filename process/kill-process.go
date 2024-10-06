package process

import (
	"fmt"
	"os"
	"os/exec"
)

func KillProcess(pids []string) error {
	args := []string{"-9"}
	for _, pid := range pids {
		args = append(args, pid)
	}

	// Logging
	file, err := os.Create("kill_output.txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	fmt.Fprintln(file, "killing", pids)
	fmt.Fprintln(file, "kill", args) // Print args to the file

	// Logging End
	// 
	cmd := exec.Command("kill", args...)

	cmd.Stdout = file
	cmd.Stderr = file
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to kill processes: %w", err)
	}

	return nil
}
