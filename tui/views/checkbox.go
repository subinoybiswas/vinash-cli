package views

import (
	"fmt"
	styles "vinash/tui/styles"
)

func checkbox(label string, checked bool, selected bool) string {
	if checked {
		return styles.CheckboxStyle.Render("[x] " + label)
	}
	if selected {
		return styles.CheckboxSelectedStyle.Render("[*] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
