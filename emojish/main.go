package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
