package models

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type ParModel struct {
	Cursor   int
	Options  []int
	Selected int
}

func NewParModel() ParModel {
	return ParModel{
		Cursor:   1,
		Options:  []int{3, 4, 5},
		Selected: 1,
	}
}

func (m ParModel) Init() tea.Cmd {
	return nil
}

func (m ParModel) Update(msg tea.Msg) (ParModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.Cursor > 0 {
				m.Cursor--
			}
		case tea.KeyDown:
			if m.Cursor < len(m.Options)-1 {
				m.Cursor++
			}
		case tea.KeyEnter:
			m.Selected = m.Options[m.Cursor]
		}
	}
	return m, nil
}

func (m ParModel) View() string {
	var s strings.Builder
	s.WriteString("Select par for this hole:\n")
	for i, option := range m.Options {
		if i == m.Cursor {
			s.WriteString(fmt.Sprintf("> %d <\n", option))
		} else {
			s.WriteString(fmt.Sprintf("  %d  \n", option))
		}
	}
	return s.String()
}
