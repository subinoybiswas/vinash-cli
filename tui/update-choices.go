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
			if m.Choice < len(m.Processess)-1 {
				m.Choice++
			}
			if m.Choice >= m.Tab*10 {
				m.Choice = m.Tab*10 - 1
			}

		case "k", "up":
			if m.Choice > 0 {
				m.Choice--
			}

		case "l", "right":
			if m.Tab < len(m.Processess)/10 {
				m.Tab++
			} else {
				m.Tab = 1
			}
			m.Choice = (m.Tab - 1) * 10

		case "h", "left":
			if m.Tab > 1 {
				m.Tab--
			} else {
				m.Tab = len(m.Processess) / 10
			}
			m.Choice = (m.Tab - 1) * 10

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
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	}

	return m, nil
}
