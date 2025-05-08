package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type HolesModel struct {
	Cursor  int
	Choices []string
	Prompt  string
	Holes   string
}

func (m HolesModel) Init() tea.Cmd {
	return nil
}

func (m HolesModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			m.Holes = m.Choices[m.Cursor]
		case "down", "j":
			m.Cursor = (m.Cursor + 1) % len(m.Choices)
		case "up", "k":
			m.Cursor = (m.Cursor - 1 + len(m.Choices)) % len(m.Choices)
		}
	}
	return m, nil
}

func (m HolesModel) View() string {
	display := m.Prompt

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		display += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	display += "\n(use arrow keys to select and enter to confirm)\n"
	return display
}
