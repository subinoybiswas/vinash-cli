package main

import (
	"fmt"
	"vinash/process"
	"vinash/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	processess, _ := process.GetProcesses()
	initialModel := tui.Model{Choice: 0, Quitting: false, Selected: []int{}, Processess: processess,Tab: 1}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
