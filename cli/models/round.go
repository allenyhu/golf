package models

import (
	"fmt"
	"strings"

	"encoding/json"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type RoundModel struct {
	Course        CourseModel `json:"course"`
	courseEntered bool        `json:"-"`
	CurrentHole   int         `json:"current_hole"`
	Holes         []HoleModel `json:"holes"`
	Done          bool        `json:"done"`
	Quitting      bool        `json:"quitting"`
}

func NewRoundModel(numHoles int) RoundModel {
	return RoundModel{
		CurrentHole:   0,
		Course:        NewCourseModel(),
		courseEntered: false,
		Done:          false,
		Quitting:      false,
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

func (m RoundModel) saveRound() {
	type holeData struct {
		Par       int    `json:"par"`
		Score     int    `json:"score"`
		TeeShot   string `json:"teeshot"`
		Putts     int    `json:"putts"`
		Chips     int    `json:"chips"`
		Penalties int    `json:"penalties"`
	}
	type roundData struct {
		Name   string     `json:"name"`
		Date   string     `json:"date"`
		Holes  int        `json:"holes"`
		Rounds []holeData `json:"hole_details"`
	}

	holes := make([]holeData, len(m.Holes))
	for i, h := range m.Holes {
		holes[i] = holeData{
			Par:       h.Par,
			Score:     h.Score,
			TeeShot:   h.TeeShot,
			Putts:     h.Putts,
			Chips:     h.Chips,
			Penalties: h.Penalties,
		}
	}

	round := roundData{
		Name:   m.Course.Name,
		Date:   m.Course.Date,
		Holes:  m.Course.Holes,
		Rounds: holes,
	}

	data, err := json.MarshalIndent(round, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling round: %v\n", err)
		return
	}

	filename := fmt.Sprintf("%s_%s.json", m.Course.Name, strings.ReplaceAll(m.Course.Date, "/", "-"))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Printf("Round saved to %s\n", filename)
}

func (m RoundModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
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
		m.saveRound()
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
