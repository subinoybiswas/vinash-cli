package views

import (
	"fmt"
	"log"
	"vinash/process"
	"vinash/tui/styles"

	"github.com/charmbracelet/lipgloss"
)

func generateIndexFrmTab(tab int) int {
	return 10 * (tab - 1)
}

func makeTabBar(tab int, total int) string {
	return fmt.Sprintf("\nPage %d/%d\n", tab, (total / 10))
}
func ChoicesView(choice int, selected []int, procesess []process.Process, tab int, width int, showDetails bool) string {
	c := choice
	finalHeading := styles.TitleStyle.Render("Vinash-CLI")
	tpl := finalHeading + "\n"
	tpl += "%s"
	tpl += makeTabBar(tab, len(procesess))
	tpl += styles.SubtleStyle.Render("j/k, up/down: select") + styles.DotStyle +
		styles.SubtleStyle.Render("enter: kill") + styles.DotStyle +
		styles.SubtleStyle.Render("h/l, left/right: tab") + styles.DotStyle +
		styles.SubtleStyle.Render("space: choose") + styles.DotStyle +
		styles.SubtleStyle.Render("q, esc: quit")

	pid := procesess[choice].Pid
	details, err := process.GetDetails(pid)
	if err != nil {
		log.Fatalf("Error getting process details: %v", err)
	}
	infoText := fmt.Sprintf(
		"PID: %s\nCommand: %s\nMemory: %.2f%%\nCPU: %.2f%%\nUID: %s\nGID: %s\nState: %s",
		details.PID,
		details.Command,
		details.Memory,
		details.CPU,
		details.UID,
		details.GID,
		details.State,
	)
	var choices string
	renderInfoText := styles.InfoStyle.Width(width/2 - 5).Render(infoText)
	choices = fmt.Sprintf(
		"%s",
		ProcessView(procesess[generateIndexFrmTab(tab):generateIndexFrmTab(tab)+10], c, selected, tab),
	)
	detailsSection := lipgloss.JoinHorizontal(lipgloss.Top, choices, renderInfoText)
	if showDetails {
		return fmt.Sprintf(tpl, detailsSection)
	} else {
		return fmt.Sprintf(tpl, choices)
	}
}
