package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Column struct {
	focused     bool
	list        list.Model
	height      int
	width       int
	foot        string
	footer      bool
	footerStyle lipgloss.Style
}

func (m Column) Init() tea.Cmd { return nil }

func (m Column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.footerStyle = lipgloss.NewStyle().Padding(1, 1).BorderStyle(lipgloss.NormalBorder()).Width(msg.Width - 2).Height(msg.Height/3 - 3).Align(lipgloss.Center)
		m.foot = m.NewFooter()
		m.width = msg.Width/margin - 2
		m.list.SetSize(msg.Width/margin, msg.Height/2)
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}
	m.list, cmd = m.list.Update(msg)
	m.foot = m.NewFooter()
	return m, cmd
}

func (m Column) View() string {
	if m.focused {
		return lipgloss.NewStyle().Padding(1, 1).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("62")).Width(m.width).Height(m.height).Render(m.list.View())
	}
	return lipgloss.NewStyle().Padding(1, 1).BorderStyle(lipgloss.NormalBorder()).Width(m.width).Height(m.height).Render(m.list.View())
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

func (m Column) NewFooter() string {
	return m.footerStyle.Render(fmt.Sprintf("%s \n\n\n\n %s", m.list.Items()[m.list.Cursor()].(Task).Title(), m.list.Items()[m.list.Cursor()].(Task).Description()))

}

func (m Column) GetFooterStyle() lipgloss.Style {
	return m.footerStyle

}
