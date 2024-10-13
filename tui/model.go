package tui

import "vinash/process"

type Model struct {
	Choice     int
	Selected   []int
	Quitting   bool
	Processess []process.Process
	Tab        int
}
