package main

import (
	"fmt"
	"vinash/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initialModel := tui.Model{Choice: 0, Quitting: false}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
