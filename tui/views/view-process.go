package views

import (
	"fmt"
	"slices"
	"vinash/process"
)

func ProcessView(processess []process.Process, checked int, selected []int) string {
	var s string
	for i, p := range processess {
		s += fmt.Sprintf("%s\n", checkbox(p.Name+" "+p.Pid, i == checked, slices.Contains(selected, i)))
	}
	return s
}
