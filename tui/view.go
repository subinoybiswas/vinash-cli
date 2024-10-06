package tui

import (
	"vinash/tui/styles"
	"vinash/tui/views"
)

func (m Model) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}

	s = views.ChoicesView(m.Choice)

	return styles.MainStyle.Render("\n" + s + "\n\n")
}
