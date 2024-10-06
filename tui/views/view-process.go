package views

import (
	"fmt"
	"vinash/process"
)

func ProcessView(processess []process.Process, checked int) string {
	var s string
	for i, p := range processess {
		s += fmt.Sprintf("%s %s\n", checkbox(p.Name, i == checked), p.Pid)
	}
	return s
}
