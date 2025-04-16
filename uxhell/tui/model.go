package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
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
	case tea.WindowSizeMsg:
		m.form = m.form.WithHeight(msg.Height - 2)
		m.form = m.form.WithWidth(msg.Width)
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

	if m.form.State == huh.StateCompleted {
		m.input += m.form.GetString("input")
		if !done {
			m.form = newForm()
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.form.State == huh.StateCompleted && done {
		return fmt.Sprintf("you wrote %s", m.input)
	}
	msg := "You haven't written anything.\n"
	if m.input != "" {
		msg = fmt.Sprintf("You've written: %s\n", m.input)
	}
	form := lipgloss.NewStyle().Render(strings.TrimSuffix(m.form.View(), "\n\n"))
	return msg + form
}
