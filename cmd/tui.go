package cmd

import (
	"fmt"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jpuriol/cuadrator/app"
	"github.com/spf13/cobra"
)

type model struct {
	data         *app.FullData
	participants []string
	cursor       int
	selected     string
	showDetails  bool
	err          error
}

var (
	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Reverse(true).
			PaddingLeft(2)

	normalStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	titleStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("230")).
			Padding(0, 1).
			Bold(true)
)

func initialModel() (*model, error) {
	d, err := app.LoadAll()
	if err != nil {
		return nil, err
	}

	var participants []string
	for name := range d.Participants {
		participants = append(participants, name)
	}
	sort.Strings(participants)

	return &model{
		data:         d,
		participants: participants,
	}, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if !m.showDetails && m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if !m.showDetails && m.cursor < len(m.participants)-1 {
				m.cursor++
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
	s.WriteString(titleStyle.Render("Participants:") + "\n\n")

	for i, name := range m.participants {
		occupations := m.data.Quadrant.GetOccupation(name)
		item := fmt.Sprintf("%s (%d occupations)", name, len(occupations))

		if m.cursor == i {
			s.WriteString(selectedStyle.Render("> "+item) + "\n")
		} else {
			s.WriteString(normalStyle.Render(item) + "\n")
		}
	}

	s.WriteString("\n(up/down: navigate, enter: details, q: quit)\n")
	return s.String()
}

func (m model) detailsView() string {
	var s strings.Builder
	s.WriteString(titleStyle.Render(fmt.Sprintf("Details for %s:", m.selected)) + "\n\n")

	occupations := m.data.Quadrant.GetOccupation(m.selected)
	if len(occupations) == 0 {
		s.WriteString("No occupations assigned.\n")
	} else {
		for _, o := range occupations {
			s.WriteString(fmt.Sprintf("- Shift %d: %s\n", o.ShiftID, m.data.Schema.OccupationName(o.OccupationID)))
		}
	}

	s.WriteString("\n(b: back, q: quit)\n")
	return s.String()
}

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Interactive TUI to view participant occupations",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := initialModel()
		if err != nil {
			return err
		}

		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
