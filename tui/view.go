package tui

import (
	"vinash/tui/styles"
	"vinash/tui/views"

	"github.com/charmbracelet/glamour"
)

func (m Model) View() string {
	var s string
	if m.Quitting {
		out, _ := glamour.Render("# Goodbye!", "dark")
		return out
	}

	s = views.ChoicesView(m.Choice, m.Selected, m.Processess, m.Tab, m.Width,m.ShowDetails)

	return styles.MainStyle.Width(m.Width).Render("\n" + s + "\n\n")
}
