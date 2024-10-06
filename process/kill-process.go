package process

import "os/exec"

func KillProcess(pid string) error {
	cmd := exec.Command("kill", pid)
	return cmd.Run()
}
