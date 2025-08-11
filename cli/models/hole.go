package models

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type HoleModel struct {
	ParModel     ParModel
	PromptCursor int
	Prompts      []string
	Prompt       string
	Inputs       []string

	Par       int
	Score     int
	TeeShot   string
	Putts     int
	Chips     int
	Penalties int

	done bool
}

func NewHoleModel() HoleModel {
	return HoleModel{
		ParModel:     NewParModel(),
		PromptCursor: 0,
		Prompts: []string{
			"What was your score on this hole?",
			"What was your tee shot result?",
			"How many putts?",
			"How many chips?",
			"How many penalties?",
		},
		Inputs: []string{"", "", "", "", ""},
	}
}

func (m HoleModel) Init() tea.Cmd {
	return nil
}

func (m HoleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			// Handle par selection first
			if m.Par == 0 {
				// Update par model
				var cmd tea.Cmd
				m.ParModel, cmd = m.ParModel.Update(msg)
				if m.ParModel.Selected > 0 {
					m.Par = m.ParModel.Selected
				}
				return m, cmd
			}

			// On enter, store the input and move to next prompt
			switch m.PromptCursor {
			case 0:
				m.Score, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[0]))
			case 1:
				m.TeeShot = strings.TrimSpace(m.Inputs[1])
			case 2:
				m.Putts, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[2]))
			case 3:
				m.Chips, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[3]))
			case 4:
				m.Penalties, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[4]))
			}
			if m.PromptCursor < len(m.Prompts)-1 {
				m.PromptCursor++
			} else {
				m.done = true
			}
		case tea.KeyBackspace, tea.KeyDelete:
			if m.Par == 0 {
				// Handle backspace in par model
				var cmd tea.Cmd
				m.ParModel, cmd = m.ParModel.Update(msg)
				return m, cmd
			}
			if len(m.Inputs[m.PromptCursor]) > 0 {
				m.Inputs[m.PromptCursor] = m.Inputs[m.PromptCursor][:len(m.Inputs[m.PromptCursor])-1]
			}
		default:
			if m.Par == 0 {
				// Handle other keys in par model
				var cmd tea.Cmd
				m.ParModel, cmd = m.ParModel.Update(msg)
				return m, cmd
			}
			// Accept only printable runes for other inputs
			if msg.Type == tea.KeyRunes {
				m.Inputs[m.PromptCursor] += msg.String()
			}
		}
	}
	return m, nil
}

func (m HoleModel) View() string {
	if m.done {
		return fmt.Sprintf(
			"Hole complete!\nPar: %d\nScore: %d\nTee Shot: %s\nPutts: %d\nChips: %d\nPenalties: %d\n",
			m.Par, m.Score, m.TeeShot, m.Putts, m.Chips, m.Penalties,
		)
	}

	if m.Par == 0 {
		return m.ParModel.View()
	}

	m.Prompt = m.Prompts[m.PromptCursor]
	input := m.Inputs[m.PromptCursor]
	return fmt.Sprintf("%s\n> %s", m.Prompt, input)
}
