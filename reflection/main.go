package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
	"github.com/charmbracelet/lipgloss"
)

const (
	fps          = 60
	spriteWidth  = 12
	spriteHeight = 5
	frequency    = 7.0
	// damping      = 0.15
	damping = 1
)

var (
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "246", Dark: "241"}).
			MarginTop(1).
			MarginLeft(2)

	spriteStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#575BD8"))
)

type frameMsg time.Time

func animate() tea.Cmd {
	return tea.Tick(time.Second/fps, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

func wait(d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return nil
	}
}

type model struct {
	x       float64
	xVel    float64
	spring  harmonica.Spring
	targetX float64
}

func (model) Init() tea.Cmd {
	return tea.Sequence(wait(time.Second/2), animate())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "left", "a", "j":
			m.targetX = 0
		case "right", "d", "k":
			m.targetX = 60
		}

	// step forward one frame
	case frameMsg:
		// update x position (and velocity) with our spring
		m.x, m.xVel = m.spring.Update(m.x, m.xVel, m.targetX)

		// return next frame
		return m, animate()

	}
	return m, nil
}

func (m model) View() string {
	var out strings.Builder
	fmt.Fprint(&out, "\n")

	x := int(math.Round(m.x))
	if x < 0 {
		return ""
	}
	mirroredX := 60 - x

	spriteRow := spriteStyle.Render(strings.Repeat("/", spriteWidth))
	row := strings.Repeat(" ", x) + spriteRow + "\n"
	fmt.Fprint(&out, strings.Repeat(row, spriteHeight))

	fmt.Fprint(&out, "\n\n")

	mirrored := strings.Repeat(" ", mirroredX) + spriteRow + "\n"
	fmt.Fprint(&out, strings.Repeat(mirrored, spriteHeight))

	fmt.Fprint(&out, helpStyle.Render("Press q to quit. Press left/right, a/d, of j/k to move left or right."))

	return out.String()
}

func main() {
	m := model{
		spring:  harmonica.NewSpring(harmonica.FPS(fps), frequency, damping),
		targetX: 60,
	}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
