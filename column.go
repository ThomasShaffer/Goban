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
	footer      *Form
	footerStyle lipgloss.Style
}

func (c Column) Init() tea.Cmd { return nil }

func (c Column) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.footerStyle = lipgloss.NewStyle().Padding(1, 1).BorderStyle(lipgloss.NormalBorder()).Width(msg.Width - 2).Height(msg.Height/3 - 3).Align(lipgloss.Center)
		c.foot = c.NewFooter()
		c.width = msg.Width/margin - 2
		c.list.SetSize(msg.Width/margin, msg.Height/2)
	case tea.KeyMsg:
		if msg.String() == "q" {
			return c, tea.Quit
		}
		if c.footer == nil || (c.footer != nil && !c.footer.active) {
			if msg.String() == "a" {
				c.footer = NewForm()
				c.foot = c.footerStyle.Render(c.footer.View())
				return c, cmd
			}
		}
		if c.footer != nil && c.footer.active {
			if msg.String() == "ctrl+b" {
				c.footer.active = false
				c.list, cmd = c.list.Update(msg)
				c.foot = c.NewFooter()
				return c, cmd
			}
			result, cmd := c.footer.Update(msg)
			c.foot = c.footerStyle.Render(result.View())
			return c, cmd
		}
	}
	c.list, cmd = c.list.Update(msg)
	c.foot = c.NewFooter()
	return c, cmd
}

func (c Column) View() string {
	if c.focused {
		return lipgloss.NewStyle().Padding(1, 1).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("62")).Width(c.width).Height(c.height).Render(c.list.View())
	}
	return lipgloss.NewStyle().Padding(1, 1).BorderStyle(lipgloss.NormalBorder()).Width(c.width).Height(c.height).Render(c.list.View())
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

func (c Column) NewFooter() string {
	return c.footerStyle.Render(fmt.Sprintf("%s \n updated: %s \n\n\n\n %s",
		c.list.Items()[c.list.Cursor()].(Task).title,
		c.list.Items()[c.list.Cursor()].(Task).date,
		c.list.Items()[c.list.Cursor()].(Task).description))

}

func (c Column) GetFooterStyle() lipgloss.Style {
	return c.footerStyle

}
