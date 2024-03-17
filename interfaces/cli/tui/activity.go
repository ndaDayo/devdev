package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type Model struct {
	focusIndex int
	inputs     [5]textinput.Model
	cursorMode cursor.Mode
	Github     *Github
}

type Github struct {
	Owner    string
	Repo     string
	Username string
	Period   Period
}

type Period struct {
	Start string
	End   string
}

func New(g *Github) *Model {
	m := &Model{
		Github: g,
	}

	for i := 0; i < 5; i++ {
		t := textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Repository owner"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Repository name"
		case 2:
			t.Placeholder = "Username"
		case 3:
			t.Placeholder = "YYYY MM DD"
		case 4:
			t.Placeholder = "YYYY MM DD"
		}

		m.inputs[i] = t
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			if msg.String() == "enter" && m.focusIndex == len(m.inputs) {
				fmt.Println("Submit!!!!")
				return m, tea.Quit
			}
			if msg.String() == "enter" && m.focusIndex < len(m.inputs) {
				switch m.focusIndex {
				case 0:
					m.Github.Owner = m.inputs[m.focusIndex].Value()
				case 1:
					m.Github.Repo = m.inputs[m.focusIndex].Value()
				case 2:
					m.Github.Username = m.inputs[m.focusIndex].Value()
				case 3:
					m.Github.Period.Start = m.inputs[m.focusIndex].Value()
				case 4:
					m.Github.Period.End = m.inputs[m.focusIndex].Value()
				}
				m.focusIndex++
				if m.focusIndex > len(m.inputs) {
					m.focusIndex = 0
				}
			} else {
				if msg.String() == "up" || msg.String() == "shift+tab" {
					m.focusIndex--
				} else {
					m.focusIndex++
				}

				if m.focusIndex > len(m.inputs) {
					m.focusIndex = 0
				} else if m.focusIndex < 0 {
					m.focusIndex = len(m.inputs)
				}
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
				} else {
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = noStyle
					m.inputs[i].TextStyle = noStyle
				}
			}

			return m, tea.Batch(cmds...)

		}
	}

	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		switch i {
		case 0:
			b.WriteString("Repository owner:\n")
		case 1:
			b.WriteString("Repository name:\n")
		case 2:
			b.WriteString("Username:\n")
		case 3:
			b.WriteString("Start (YYYY MM DD):\n")
		case 4:
			b.WriteString("End (YYYY MM DD):\n")
		}

		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
