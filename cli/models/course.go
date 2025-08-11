package models

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type CourseModel struct {
	Name  string
	Date  string
	Holes int

	Prompts      []string
	Inputs       []string
	PromptCursor int

	done bool
}

func NewCourseModel() CourseModel {
	return CourseModel{
		Prompts: []string{
			"Course Name: ",
			"Date (MM/DD/YYYY): ",
			"Number of Holes Played: ",
		},
		Inputs:       []string{"", "", ""},
		PromptCursor: 0,
		done:         false,
	}
}

func (m CourseModel) Init() tea.Cmd {
	return nil
}

func (m CourseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			switch m.PromptCursor {
			case 0:
				fmt.Printf("Name: %s", m.Inputs[0])
				m.Name = m.Inputs[0]
			case 1:
				m.Date = m.Inputs[1]
			case 2:
				// Prob can take msg as int
				m.Holes, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[2]))
			}

			if m.PromptCursor < len(m.Prompts)-1 {
				m.PromptCursor++
			} else {
				m.done = true
				return m, nil
			}
		case tea.KeyBackspace, tea.KeyDelete:
			if len(m.Inputs[m.PromptCursor]) > 0 {
				m.Inputs[m.PromptCursor] = m.Inputs[m.PromptCursor][:len(m.Inputs[m.PromptCursor])-1]
			}
		default:
			if msg.Type == tea.KeyRunes {
				m.Inputs[m.PromptCursor] += msg.String()
			}
		}
	}

	return m, nil
}

func (m CourseModel) View() string {
	return fmt.Sprintf("%s\n> %s", m.Prompts[m.PromptCursor], m.Inputs[m.PromptCursor])
}
