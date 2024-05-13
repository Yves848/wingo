package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gopkg.in/toast.v1"
)

type model struct {
	textInput   textinput.Model
	choices     []string
	cursor      int
	selected    map[int]struct{}
	PromptStyle lipgloss.Style
}

func initialModel() model {
	// PromptStyle.Width(30)
	t1 := textinput.New()
	t1.Placeholder = "Type here"
	t1.Prompt = "⟩ "
	t1.TextStyle = lipgloss.NewStyle().Background(lipgloss.Color("#000000")).Foreground(lipgloss.Color("#FFFFFF"))
	t1.Width = 30
	t1.Focus()

	return model{
		textInput: t1,
		choices:   []string{"Choix 1", "Choix 2", "Choix 3"},
		selected:  make(map[int]struct{}),
		cursor:    0,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, textinput.Blink)
	// return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyDown:
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		}

	}
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	PromptStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	PromptStyle.Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FF00FF"))
	s := PromptStyle.Render((m.textInput.View()))
	s += "\n\n"

	for i, choice := range m.choices {
		if i == m.cursor {
			s += fmt.Sprintf("-> %s\n", choice)
		} else {
			s += fmt.Sprintf("   %s\n", choice)
		}
	}

	s += "\n\nEsc to quit"
	return s
}

func main() {
	notification := toast.Notification{
		AppID:   "Example App",
		Title:   "My notification",
		Message: "Some message about how important something is...",
		// Icon:    "go.png", // This file must exist (remove this line if it doesn't)
		Actions: []toast.Action{
			{Type: "protocol", Label: "I'm a button", Arguments: ""},
			{Type: "protocol", Label: "Me too!", Arguments: ""},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
	}
}
