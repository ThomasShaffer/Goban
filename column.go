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

func renderColumns(data []todoModel) []Column {
	var columnList []Column
	var column Column
	var todoData todoModel

	todoData = data[0]

	var delegate list.DefaultDelegate

	var projectModel = [][]map[string]string{todoData.todo, todoData.doing, todoData.did}
	for categoryIndex := range projectModel {
		var categoryList list.Model
		var cat status
		delegate = list.NewDefaultDelegate()
		categoryList = list.New([]list.Item{}, delegate, 0, 0)
		categoryList.SetShowHelp(false)
		switch categoryIndex {
		case 0:
			categoryList.Title = "todo"
			cat = todo
		case 1:
			categoryList.Title = "doing"
			cat = doing
		case 2:
			categoryList.Title = "did"
			cat = did
		}
		category := projectModel[categoryIndex]
		var itemList []list.Item
		for itemKey := range category {
			var item list.Item
			item = Task{
				status:      cat,
				title:       category[itemKey]["title"],
				description: category[itemKey]["description"],
				date:        category[itemKey]["date"],
			}
			itemList = append(itemList, item)
		}
		categoryList.SetItems(itemList)

		column = Column{focused: true, list: categoryList, width: 30, height: 30}
		columnList = append(columnList, column)
	}

	return columnList
}

func (m Column) NewFooter() string {
	return m.footerStyle.Render(fmt.Sprintf("%s \n updated: %s \n\n\n\n %s",
		m.list.Items()[m.list.Cursor()].(Task).title,
		m.list.Items()[m.list.Cursor()].(Task).date,
		m.list.Items()[m.list.Cursor()].(Task).description))

}

func (m Column) GetFooterStyle() lipgloss.Style {
	return m.footerStyle

}
