package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	. "github.com/wederbrand/advent-of-code/github.com/wederbrand/chart"
	"os"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Italic(true).
			Foreground(lipgloss.Color("#FFD700")). // Gold
			Background(lipgloss.Color("#333333"))  // Dark gray background

	counterStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#32CD32")). // Lime green
			Background(lipgloss.Color("#000000"))  // Black background

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ADD8E6")). // Light blue
			Italic(true)
)

type Model struct {
	c Chart
}

func New(c Chart) {
	p := tea.NewProgram(Model{})
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app: %v", err)
		os.Exit(1)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	title := titleStyle.Render("Stylish Counter")
	counter := counterStyle.Render(fmt.Sprintf("Counter: %d", m.counter))

	return fmt.Sprintf("%s\n\n%s%s", title, counter)
}
