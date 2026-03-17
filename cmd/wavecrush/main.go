package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles
var (
	titleStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00D9FF")).
		MarginTop(1).
		MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888"))

	menuItemStyle = lipgloss.NewStyle().
		PaddingLeft(2)

	selectedItemStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00D9FF")).
		Bold(true).
		PaddingLeft(2)
)

type model struct {
	cursor   int
	options  []string
	selected map[int]struct{}
}

func initialModel() model {
	return model{
		options: []string{
			"Convert Audio Format",
			"Batch Processing",
			"Audio Effects",
			"Metadata Editor",
			"Waveform Visualization",
			"Settings",
			"Quit",
		},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "enter", " ":
			if m.cursor == len(m.options)-1 {
				return m, tea.Quit
			}
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("🌊 wavecrush") + "\n"
	s += subtitleStyle.Render("Beautiful FFmpeg TUI - Built with Charm") + "\n\n"

	for i, option := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		if m.cursor == i {
			s += selectedItemStyle.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, option)) + "\n"
		} else {
			s += menuItemStyle.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, option)) + "\n"
		}
	}

	s += "\n" + subtitleStyle.Render("Press j/k or ↑/↓ to navigate, enter to select, q to quit") + "\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
