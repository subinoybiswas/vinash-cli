package views

import (
	"fmt"
	styles "vinash/tui/styles"
)

func checkbox(label string, checked bool) string {
	if checked {
		return styles.CheckboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
