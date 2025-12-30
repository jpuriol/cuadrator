package tui

import (
	"fmt"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jpuriol/cuadrator/core"
)

type model struct {
	participants []string
	cursor       int
	selected     string
	showDetails  bool
	err          error

	// Cached data
	schedule core.Schedule
	schema   core.Schema

	height     int
	startIndex int
}

func NewModel(participants core.Participants, schedule core.Schedule, schema core.Schema) tea.Model {
	var pList []string
	for name := range participants {
		pList = append(pList, name)
	}
	sort.Strings(pList)

	return &model{
		participants: pList,
		schedule:     schedule,
		schema:       schema,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if !m.showDetails && m.cursor > 0 {
				m.cursor--
				if m.cursor < m.startIndex {
					m.startIndex = m.cursor
				}
			}

		case "down", "j":
			if !m.showDetails && m.cursor < len(m.participants)-1 {
				m.cursor++
				maxVisible := m.height - 5
				if m.height == 0 {
					maxVisible = len(m.participants)
				} else if maxVisible < 1 {
					maxVisible = 1
				}
				if m.cursor >= m.startIndex+maxVisible {
					m.startIndex = m.cursor - maxVisible + 1
				}
			}

		case "enter":
			if !m.showDetails {
				m.selected = m.participants[m.cursor]
				m.showDetails = true
			}

		case "esc", "backspace", "b":
			if m.showDetails {
				m.showDetails = false
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}

	if m.showDetails {
		return m.detailsView()
	}

	return m.listView()
}

func (m model) listView() string {
	var s strings.Builder
	s.WriteString("PARTICIPANTS:\n\n")

	maxVisible := m.height - 5
	if m.height == 0 {
		maxVisible = len(m.participants)
	} else if maxVisible < 1 {
		maxVisible = 1
	}

	endIndex := m.startIndex + maxVisible
	if endIndex > len(m.participants) {
		endIndex = len(m.participants)
	}

	for i := m.startIndex; i < endIndex; i++ {
		name := m.participants[i]
		assignments := m.schedule.GetAssignments(name)
		item := fmt.Sprintf("%s (%d occupations)", name, len(assignments))

		if m.cursor == i {
			s.WriteString(fmt.Sprintf("> %s <\n", item))
		} else {
			s.WriteString(fmt.Sprintf("  %s  \n", item))
		}
	}

	// Add padding if necessary to keep the footer at the bottom if we want,
	// but simple scrolling might be enough.
	// The original code didn't have padding.

	s.WriteString("\n(up/down: navigate, enter: details, q: quit)\n")
	return s.String()
}

func (m model) detailsView() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("Details for %s:\n\n", m.selected))

	assignments := m.schedule.GetAssignments(m.selected)
	if len(assignments) == 0 {
		s.WriteString("No occupations assigned.\n")
	} else {
		for _, o := range assignments {
			s.WriteString(fmt.Sprintf("- Shift %d: %s\n", o.ShiftID, m.schema.OccupationName(o.OccupationID)))
		}
	}

	s.WriteString("\n(b: back, q: quit)\n")
	return s.String()
}
