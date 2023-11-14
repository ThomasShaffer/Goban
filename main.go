package main

import (
	"fmt"
	flexbox "github.com/76creates/stickers"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type gobanModel struct {
	flexBox  *flexbox.FlexBox
	todos    []string
	cursor   int
	selected map[int]struct{}
}

var (
	style1 = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Align(lipgloss.Center)
	style2 = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Align(lipgloss.Center)
	style3 = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Align(lipgloss.Center)
	style4 = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Align(lipgloss.Center)
	style5 = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Align(lipgloss.Center)
)

func (m *gobanModel) Init() tea.Cmd { return nil }

func (m *gobanModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.flexBox.SetWidth(msg.Width)
		m.flexBox.SetHeight(msg.Height)
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		case "k":
			if m.cursor > 0 {
				m.cursor--
			}
			return m, nil
		case "j":
			if m.cursor < len(m.todos) {
				m.cursor++
			}
			return m, nil
		}
	}
	return m, nil
}

func (m *gobanModel) View() string {
	row, _ := m.flexBox.GetRow(0)
	row.Cell(0).SetContent("goban")

	secondRow, _ := m.flexBox.GetRow(1)
	secondRow.Cell(0).SetContent("todo")
	secondRow.Cell(1).SetContent("doing")
	secondRow.Cell(2).SetContent("did")

	thirdRow, _ := m.flexBox.GetRow(2)
	thirdRow.Cell(0).SetContent("footer")

	return m.flexBox.Render()
}

func main() {

	m := gobanModel{
		flexBox: flexbox.NewFlexBox(30, 30),
	}

	rows := []*flexbox.FlexBoxRow{
		m.flexBox.NewRow().AddCells([]*flexbox.FlexBoxCell{
			flexbox.NewFlexBoxCell(1, 1).SetStyle(style1),
		}),
		m.flexBox.NewRow().AddCells([]*flexbox.FlexBoxCell{
			flexbox.NewFlexBoxCell(2, 10).SetStyle(style2),
			flexbox.NewFlexBoxCell(5, 10).SetStyle(style3),
			flexbox.NewFlexBoxCell(2, 10).SetStyle(style4),
		}),
		m.flexBox.NewRow().AddCells([]*flexbox.FlexBoxCell{
			flexbox.NewFlexBoxCell(1, 1).SetStyle(style5),
		}),
	}

	m.flexBox.AddRows(rows)

	p := tea.NewProgram(&m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
