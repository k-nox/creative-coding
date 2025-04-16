package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const (
	letters     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	punctuation = "!@#$%^&*()_+-=[]{}/\\|,.;'<>?:\""
)

var done bool

type model struct {
	form  *huh.Form
	input string
}

func New() model {
	return model{
		form: newForm(),
	}
}

func (m model) Init() tea.Cmd {
	return m.form.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	}
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return m, cmd
}

func (m model) View() string {
	if m.form.State == huh.StateCompleted {
		return fmt.Sprintf("you chose %s", m.form.GetString("input"))
	}
	return m.form.View()
}
