package tui

import (
	"slices"
	"time"
	"vinash/process"

	tea "github.com/charmbracelet/bubbletea"
)

func updateChoices(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 9 {
				m.Choice = 0
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case " ":
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
		case "enter":
			if len(m.Selected) > 0 {
				toBeKilled := []string{}
				for _, choice := range m.Selected {
					pid := m.Processess[choice].Pid
					toBeKilled = append(toBeKilled, pid)
				}
				process.KillProcess(toBeKilled)
			}
			m.Quitting = true
			return m, tea.Quit
		}
	case tickMsg:
		processess, _ := process.GetProcesses()
		m.Processess = processess
		return m, tickAfter(time.Second)

	}

	return m, nil
}
