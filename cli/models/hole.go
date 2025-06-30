package models

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type HoleModel struct {
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
		PromptCursor: 0,
		Prompts: []string{
			"What is the par for this hole?",
			"What was your score on this hole?",
			"What was your tee shot result?",
			"How many putts?",
			"How many chips?",
			"How many penalties?",
		},
		Inputs: []string{"", "", "", "", "", ""},
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
			// On enter, store the input and move to next prompt
			switch m.PromptCursor {
			case 0:
				m.Par, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[0]))
			case 1:
				m.Score, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[1]))
			case 2:
				m.TeeShot = strings.TrimSpace(m.Inputs[2])
			case 3:
				m.Putts, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[3]))
			case 4:
				m.Chips, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[4]))
			case 5:
				m.Penalties, _ = strconv.Atoi(strings.TrimSpace(m.Inputs[5]))
			}
			if m.PromptCursor < len(m.Prompts)-1 {
				m.PromptCursor++
			} else {
				m.done = true
			}
		case tea.KeyBackspace, tea.KeyDelete:
			if len(m.Inputs[m.PromptCursor]) > 0 {
				m.Inputs[m.PromptCursor] = m.Inputs[m.PromptCursor][:len(m.Inputs[m.PromptCursor])-1]
			}
		default:
			// Accept only printable runes
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
	m.Prompt = m.Prompts[m.PromptCursor]
	input := m.Inputs[m.PromptCursor]
	return fmt.Sprintf("%s\n> %s", m.Prompt, input)
}
