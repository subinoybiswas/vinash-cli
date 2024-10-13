package views

import (
	"fmt"
	"vinash/process"
	"vinash/tui/styles"

	"github.com/charmbracelet/glamour"
)

func generateIndexFrmTab(tab int) int {
	return 10 * (tab - 1)
}

func makeTabBar(tab int, total int) string {
	return fmt.Sprintf("\nPage %d/%d\n", tab, (total/10))
}
func ChoicesView(choice int, selected []int, procesess []process.Process, tab int) string {
	c := choice
	heading,err:=glamour.Render("# Choose processes to kill", "dark")
	if err!=nil{
		fmt.Println(err)
	}
	tpl :=heading
	tpl += "%s"
	tpl += makeTabBar(tab, len(procesess))
	tpl += styles.SubtleStyle.Render("j/k, up/down: select") + styles.DotStyle +
		styles.SubtleStyle.Render("enter: kill") + styles.DotStyle +
		styles.SubtleStyle.Render("h/l, left/right: tab") + styles.DotStyle +
		styles.SubtleStyle.Render("space: choose") + styles.DotStyle +
		styles.SubtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s",
		ProcessView(procesess[generateIndexFrmTab(tab):generateIndexFrmTab(tab)+10], c, selected),
	)

	return fmt.Sprintf(tpl, choices)
}
