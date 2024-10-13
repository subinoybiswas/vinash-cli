package styles

import "github.com/charmbracelet/lipgloss"

const (
	dotChar           = " • "
)


var (
	SubtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	CheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	CheckboxSelectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("112"))
	DotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	MenuDotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
	MainStyle     = lipgloss.NewStyle().MarginLeft(2)
)
