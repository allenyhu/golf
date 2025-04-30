package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// Model represents the program's state
type model struct {
	cursor   int
	choices  []string
	selected bool
	quitting bool
	holes    int
	prompt   string
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
			if m.choices[m.cursor] == "Yes" {
				m.selected = true
			} else {
				m.quitting = true
				return m, tea.Quit
			}
		case "down", "j":
			m.cursor = (m.cursor + 1) % len(m.choices)
		case "up", "k":
			m.cursor = (m.cursor - 1 + len(m.choices)) % len(m.choices)
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Quitting...\n"
	}
	if m.selected {
		choice := m.choices[m.cursor]
		return fmt.Sprintf("You chose: %s\n", choice)
	}

	s := m.prompt

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
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
		p := tea.NewProgram(initialModel())
		m, err := p.Run()
		if err != nil {
			fmt.Printf("Error running program: %v", err)
			return
		}

		// Type assert the final model to access its fields
		finalModel := m.(model)
		if !finalModel.quitting && finalModel.selected && finalModel.cursor == 0 {
			fmt.Println("Starting golf score tracking...")
			// Add your score tracking logic here
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
