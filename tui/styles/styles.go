package styles

import "github.com/charmbracelet/lipgloss"

const (
	dotChar           = " • "
)


var (
	SubtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	CheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	DotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	MainStyle     = lipgloss.NewStyle().MarginLeft(2)
)
