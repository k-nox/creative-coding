package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/k-nox/creative-coding/uxhell/tui"
)

func main() {
	if _, err := tea.NewProgram(tui.New(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal("Oops:", err)
	}
}
