package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
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
		c.footerStyle = lipgloss.NewStyle().Padding(1, 1).BorderStyle(config.borderStyle).Width(msg.Width - 2).Height(msg.Height/3 - 3).Align(lipgloss.Center)
		c.foot = c.NewFooter()
		c.width = msg.Width/margin - 2
		c.list.SetSize(msg.Width/margin, msg.Height/2)
	case tea.KeyMsg:
		if c.footer == nil || (c.footer != nil && !c.footer.active) {
			switch msg.String() {
			case "a":
				c.footer = NewForm()
				c.foot = c.footerStyle.Render(c.footer.View())
				return c, cmd
			case "e":
				c.footer = EditForm(c.list.SelectedItem().(Task))
				c.foot = c.footerStyle.Render(c.footer.View())
				return c, cmd
			case "D":
				DeleteDataInJson(c.list.SelectedItem().(Task))
				c.list.RemoveItem(c.list.Cursor())
				c.list.ResetSelected()
				return c, cmd
			case "J":
				currentIndex, nextIndex, taskStatus := swapItems(&c, 1)
				MoveDataInJson(currentIndex, nextIndex, taskStatus)
			case "K":
				currentIndex, nextIndex, taskStatus := swapItems(&c, -1)
				MoveDataInJson(currentIndex, nextIndex, taskStatus)
			case "P":
				c.footer = PomodoroForm()
				c.foot = c.footerStyle.Render(c.footer.View())
				return c, cmd
			case "N":
				c.footer = ProjectForm()
				c.foot = c.footerStyle.Render(c.footer.View())
				return c, cmd
			}
		}
		if c.footer != nil && c.footer.active {
			switch msg.String() {
			case "ctrl+b":
				c.footer.active = false
				c.list, cmd = c.list.Update(msg)
				c.foot = c.NewFooter()
				return c, cmd
			case "enter":
				currTime := time.Now()
				userTask := Task{
					status:      todo,
					title:       c.footer.title.Value(),
					description: c.footer.description.Value(),
					date:        currTime.Format("01-01-2006"),
				}
				if c.footer.formType == edit {
					EditDataInJson(userTask, c.list.SelectedItem().(Task))
					c.footer.active = false
					c.list.SetItem(c.list.Cursor(), userTask)

				} else if c.footer.formType == pomodoro {
					c.footer.active = false

				} else if c.footer.formType == add {
					AddDataToJson(userTask)
					c.footer.active = false
					c.list.InsertItem(100, userTask)
				} else if c.footer.formType == project {
					WriteNewJson(userTask)
					c.footer.active = false
				}
			default:
				result, cmd := c.footer.Update(msg)
				c.foot = c.footerStyle.Render(result.View())
				return c, cmd
			}
		}
	}
	c.list, cmd = c.list.Update(msg)
	c.foot = c.NewFooter()
	return c, cmd
}

func (c Column) View() string {
	if c.focused {
		return lipgloss.NewStyle().Padding(1, 1).BorderStyle(config.borderStyle).BorderForeground(lipgloss.Color("62")).Width(c.width).Height(c.height - config.columnHeight).Render(c.list.View())
	}
	return lipgloss.NewStyle().Padding(1, 1).BorderStyle(config.borderStyle).Width(c.width).Height(c.height - 15).Render(c.list.View())
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

func swapItems(c *Column, dir int) (int, int, string) {
	currentTask := c.list.SelectedItem()
	currentIndex := c.list.Cursor()
	if dir == 1 {
		c.list.CursorDown()
	} else if dir == -1 {
		c.list.CursorUp()
	}
	c.list.Select(c.list.Cursor())
	nextTask := c.list.SelectedItem()
	nextIndex := c.list.Cursor()
	c.list.SetItem(nextIndex, currentTask)
	c.list.SetItem(currentIndex, nextTask)
	return currentIndex, nextIndex, StatusToString(currentTask.(Task).status)
}
