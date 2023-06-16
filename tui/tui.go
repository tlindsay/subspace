package tui

import (
	"fmt"
	"os"

	"github.com/tlindsay/subspace/subspace"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type list struct {
	selected bool
	name     string
}
type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	table    table.Model
}

func initialModel() model {
	columns := []table.Column{
		{Title: "", Width: 3},
		{Title: "Character", Width: 10},
		{Title: "File", Width: 10},
	}
	rows := []table.Row{}
	for _, c := range subspace.ListAllCharacters() {
		rows = append(rows, table.Row{c, fmt.Sprintf("%s.txt", c)})
	}

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	return model{
		choices:  subspace.ListAllCharacters(),
		selected: make(map[int]struct{}),
		table: table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithHeight(12),
			table.WithStyles(s),
		),
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
			return m, tea.Quit
		case "up", "k":
			cursor := m.table.Cursor()
			if cursor > 0 {
				m.table.SetCursor(cursor - 1)
			}
		case "down", "j":
			cursor := m.table.Cursor()
			if cursor < len(m.choices) {
				m.table.SetCursor(cursor + 1)
			}
		case "enter", " ":
			return m, tea.Batch(
				tea.Printf("Selected %s", m.table.SelectedRow()[0]),
			)
		}
	}

	return m, nil
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func StartTUI() {
	main()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Red alert! The warp core has been breached: %v", err)
		os.Exit(1)
	}
}
