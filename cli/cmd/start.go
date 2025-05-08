package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"cli/models"
)

// Model represents the program's state
type model struct {
	choicesCursor     int
	choices           []string
	selected          bool
	quitting          bool
	holes             int
	prompt            string
	previousSelection string
	modelsCursor      int
	models            []tea.Model

	// need a state that toggles which model should be used
	// then send message to the model
	// https://www.youtube.com/watch?v=uJ2egAkSkjg
}

// Initial model
func initialModel() model {
	return model{
		choices: []string{"Yes", "No"},
		prompt:  "Would you like to start tracking your golf score?\n\n",
		models: []tea.Model{
			models.HolesModel{
				Choices: []string{"9", "18"},
				Prompt:  "How many holes did you play?\n\n",
			},
			models.HoleModel{},
		},
		modelsCursor: -1,
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
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.choices[m.choicesCursor] == "Yes" {
				m.selected = true
				m.previousSelection = m.choices[m.choicesCursor]

				m.modelsCursor++
				newModel := m.models[m.modelsCursor]

				// need to fix model progression now
				m.prompt = newModel.Prompt

				// choice is not always going to be preset for Model
				// is there a way to release control to the child model
				m.choicesCursor = 0
				m.choices = newModel.Choices

				return newModel, nil
			} else {
				m.quitting = true
				return m, tea.Quit
			}
		case "down", "j":
			m.choicesCursor = (m.choicesCursor + 1) % len(m.choices)
		case "up", "k":
			m.choicesCursor = (m.choicesCursor - 1 + len(m.choices)) % len(m.choices)
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Quitting...\n"
	}
	s := ""
	if m.selected {
		s += fmt.Sprintf("You chose: %s\n", m.previousSelection)

	}

	s += m.prompt

	for i, choice := range m.choices {
		cursor := " "
		if m.choicesCursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	s += "\n(use arrow keys to select and enter to confirm)\n"
	return s
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking golf score",
	Run: func(cmd *cobra.Command, args []string) {
		// p := tea.NewProgram(initialModel())
		p := tea.NewProgram(models.NewHoleModel())
		m, err := p.Run()
		if err != nil {
			fmt.Printf("Error running program: %v", err)
			return
		}

		// Type assert the final model to access its fields
		finalModel := m.(model)
		if !finalModel.quitting && finalModel.selected && finalModel.choicesCursor == 0 {
			fmt.Println("Starting golf score tracking...")
			// Add your score tracking logic here
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
