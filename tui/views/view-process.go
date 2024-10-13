package views

import (
	"fmt"
	"slices"
	"vinash/process"
)

func ProcessView(processess []process.Process, checked int, selected []int, tab int) string {
	var s string
	for i, p := range processess {
		s += fmt.Sprintf("%s\n", checkbox(p.Name+" "+p.Pid, i+10*(tab-1) == checked, slices.Contains(selected, i+10*(tab-1))))
	}
	return s
}
