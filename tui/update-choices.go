package tui

import (
	"slices"

	tea "github.com/charmbracelet/bubbletea"
)

func updateChoices(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 19 {
				m.Choice = 0
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			if slices.Contains(m.Selected, m.Choice) {
				for i, choice := range m.Selected {
					if choice == m.Choice {
						m.Selected = append(m.Selected[:i], m.Selected[i+1:]...)
						break
					}
				}
			} else {
				m.Selected = append(m.Selected, m.Choice)
			}

		}

	}

	return m, nil
}
