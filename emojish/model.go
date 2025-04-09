package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	inp textinput.Model
}

type execMsg struct {
	out string
}

func newModel() model {
	inp := textinput.New()
	inp.Focus()

	return model{
		inp: inp,
	}
}

// Init implements tea.Model.
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements tea.Model.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			return m, m.exec
		}
	case execMsg:
		out := msg.out
		history := m.replacedView()
		m.inp.SetValue("")
		return m, tea.Println(history, "\n", out)
	}

	m.inp, cmd = m.inp.Update(msg)
	return m, cmd
}

// View implements tea.Model.
func (m model) View() string {
	return m.replacedView()
}

func (m model) exec() tea.Msg {
	raw := strings.Split(m.inp.Value(), " ")
	if len(raw) == 0 {
		return nil
	}

	cmd := raw[0]
	var args []string
	if len(raw) > 1 {
		args = raw[1:]
	}

	c := exec.Command(cmd, args...)
	out, err := c.CombinedOutput()
	if err != nil {
		return execMsg{
			out: fmt.Sprintf("%s: %s", err.Error(), out),
		}
	}

	return execMsg{
		out: string(out),
	}
}

func (m model) replacedView() string {
	actual := m.inp.Value()
	replacements := []string{}
	for letter, emoji := range emojis {
		replacements = append(replacements, letter, emoji)
	}

	replacer := strings.NewReplacer(replacements...)
	mapped := replacer.Replace(actual)
	m.inp.SetValue(mapped)
	out := m.inp.View()
	m.inp.SetValue(actual)
	return out
}
