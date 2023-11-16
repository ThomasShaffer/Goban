package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Column struct {
	focused bool
	list    list.Model
	height  int
	width   int
}

func (m Column) Init() tea.Cmd { return nil }

func (m Column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width/margin - 2
		m.list.SetSize(msg.Width/margin, msg.Height/2)
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "k":
			m.list.CursorUp()
			return m, nil
		case "j":
			m.list.CursorDown()
			return m, nil
		}
	}
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Column) View() string {
	if m.focused {
		return lipgloss.NewStyle().Padding(1, 2).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("62")).Width(m.width).Height(m.height).Render(m.list.View())
	}
	return lipgloss.NewStyle().Padding(1, 2).BorderStyle(lipgloss.NormalBorder()).Width(m.width).Height(m.height).Render(m.list.View())
}

func NewColumn(title string, width, height int) Column {
	delegate := list.NewDefaultDelegate()
	newList := list.New([]list.Item{}, delegate, 0, 0)
	newList.SetShowHelp(false)
	newList.Title = title
	//newList.SetShowTitle(true)

	newList.SetItems([]list.Item{
		Task{status: todo, title: "do this", description: "now"},
		Task{status: todo, title: "do this 1", description: "ok"},
		Task{status: todo, title: "do this 2", description: "yerr"},
	})
	return Column{focused: true, list: newList, height: height, width: width}
}

func NewBottomColumn(width, height int) Column {
	delegate := list.NewDefaultDelegate()
	newList := list.New([]list.Item{}, delegate, 0, 0)
	newList.SetShowHelp(false)
	newList.Title = "description"
	//newList.SetShowTitle(true)

	newList.SetItems([]list.Item{
		Task{status: todo, title: "do this", description: "now"},
	})
	return Column{focused: true, list: newList, height: height, width: width}

}
