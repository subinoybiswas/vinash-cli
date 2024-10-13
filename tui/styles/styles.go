package styles

import (
	"image/color"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

const (
	dotChar = " • "
)

var (
	SubtleStyle           = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	CheckboxStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	CheckboxSelectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("112"))
	DotStyle              = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	MenuDotStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))
	MainStyle             = lipgloss.NewStyle().MarginLeft(2).MarginRight(2)
	TitleStyle            = lipgloss.NewStyle().
				MarginLeft(1).
				MarginRight(1).
				MarginBottom(1).
				
				Padding(0, 1).
				Italic(true).
				Foreground(lipgloss.Color("#FFF7DB")).Background(lipgloss.Color("#F25D94"))
)

func Rainbow(base lipgloss.Style, s string, colors []color.Color) string {
	var str string
	for i, ss := range s {
		color, _ := colorful.MakeColor(colors[i%len(colors)])
		str = str + base.Background(lipgloss.Color(color.Hex())).Render(string(ss))
	}
	return str
}
