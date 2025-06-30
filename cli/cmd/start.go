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
	prompt            string
	previousSelection string
}

// Initial model
func initialModel() model {
	return model{
		choices: []string{"Yes", "No"},
		prompt:  "Would you like to start tracking your golf score?\n\n",
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
				return m, nil
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

		// Start the round with 18 holes
		roundModel := models.NewRoundModel(18)
		p := tea.NewProgram(roundModel)
		model, err := p.Run()

		if err != nil {
			fmt.Printf("Error running program: %v", err)
			return
		}

		// The round is complete, finalModel contains all hole data
		finalModel := model.(models.RoundModel)
		if finalModel.Done {
			fmt.Println("Round complete! Thank you for tracking your score.")
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
