package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type RoundModel struct {
	Course        CourseModel
	courseEntered bool

	CurrentHole int
	Holes       []HoleModel
	Done        bool
	Quitting    bool
}

func NewRoundModel(numHoles int) RoundModel {
	// holes := make([]HoleModel, numHoles)
	// for i := range holes {
	// 	holes[i] = NewHoleModel()
	// }

	return RoundModel{
		CurrentHole:   0,
		Course:        NewCourseModel(),
		courseEntered: false,
		// Holes:       holes,
		Done:     false,
		Quitting: false,
	}
}

func (m RoundModel) Init() tea.Cmd {
	return nil
}

func (m RoundModel) generateHoles(numHoles int) []HoleModel {
	holes := make([]HoleModel, numHoles)
	for i := range holes {
		holes[i] = NewHoleModel()
	}

	return holes
}

func (m RoundModel) generateSummary() string {
	summary := "Round Summary\n\n"
	totalScore := 0
	totalPar := 0
	holesPlayed := 0

	for i, hole := range m.Holes {
		// Only include holes that have been played (have a score)
		if hole.Score > 0 {
			holesPlayed++
			summary += fmt.Sprintf("Hole %d:\n", i+1)
			summary += fmt.Sprintf("  Par: %d\n", hole.Par)
			summary += fmt.Sprintf("  Score: %d\n", hole.Score)
			summary += fmt.Sprintf("  Tee Shot: %s\n", hole.TeeShot)
			summary += fmt.Sprintf("  Putts: %d\n", hole.Putts)
			summary += fmt.Sprintf("  Chips: %d\n", hole.Chips)
			summary += fmt.Sprintf("  Penalties: %d\n\n", hole.Penalties)

			totalScore += hole.Score
			totalPar += hole.Par
		}
	}

	if holesPlayed > 0 {
		summary += fmt.Sprintf("Holes Played: %d of %d\n", holesPlayed, m.Course.Holes)
		summary += fmt.Sprintf("Total Score: %d\n", totalScore)
		summary += fmt.Sprintf("Total Par: %d\n", totalPar)
		summary += fmt.Sprintf("Score vs Par: %+d\n", totalScore-totalPar)
	} else {
		summary += "No holes were played.\n"
	}

	return summary
}

func (m RoundModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Quitting = true
			return m, tea.Quit
		}

		if !m.courseEntered {
			updatedCourse, cmd := m.Course.Update(msg)
			if courseModel, ok := updatedCourse.(CourseModel); ok {
				m.Course = courseModel
				if m.Course.done {
					m.courseEntered = true
					m.Holes = m.generateHoles(m.Course.Holes)
				}
			}

			return m, cmd
		}

		// Update the current hole model
		currentHole := m.Holes[m.CurrentHole]
		updatedHole, cmd := currentHole.Update(msg)

		// Type assert back to HoleModel since Update returns tea.Model
		if holeModel, ok := updatedHole.(HoleModel); ok {
			m.Holes[m.CurrentHole] = holeModel

			// If the current hole is done, move to the next hole
			if holeModel.done {
				if m.CurrentHole < m.Course.Holes-1 {
					m.CurrentHole++
				} else {
					m.Done = true
				}
			}
		}

		return m, cmd
	}
	return m, nil
}

func (m RoundModel) View() string {
	if m.Done || m.Quitting {
		return m.generateSummary()
	}

	if !m.courseEntered {
		return m.Course.View()
	}

	// Show current hole number and its view
	currentHole := m.Holes[m.CurrentHole]
	return fmt.Sprintf("Hole %d of %d\n\n%s",
		m.CurrentHole+1,
		m.Course.Holes,
		currentHole.View())
}
