package views

import (
	"fmt"
	"vinash/process"
	"vinash/tui/styles"
)

func ChoicesView(choice int, selected []int) string {
	c := choice

	tpl := "Select Process to kill\n\n"
	tpl += "%s\n\n"
	tpl += styles.SubtleStyle.Render("j/k, up/down: select") + styles.DotStyle +
		styles.SubtleStyle.Render("space: choose") + styles.DotStyle +
		styles.SubtleStyle.Render("q, esc: quit")
	procesess, _ := process.GetProcesses()
	procesess = procesess[:10]
	choices := fmt.Sprintf(
		"%s\n",
		ProcessView(procesess, c,selected),
	)

	return fmt.Sprintf(tpl, choices)
}
