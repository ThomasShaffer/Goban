package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListModel struct {
	focused status
	lists   []Column
	header  string
	loaded  bool
}

func (m *ListModel) Init() tea.Cmd { return nil }
func (m *ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		var cmds []tea.Cmd
		m.header = lipgloss.NewStyle().Padding(1, 1).BorderStyle(lipgloss.NormalBorder()).Width(msg.Width - 2).Height(2).Align(lipgloss.Center).Render("\nGoban")
		m.lists = m.initializeLists()
		for i := 0; i < len(m.lists); i++ {
			var res tea.Model
			res, cmd = m.lists[i].Update(msg)
			m.lists[i] = res.(Column)
			if i == 0 {
				m.lists[i].focused = true
			} else {
				m.lists[i].focused = false
			}
			cmds = append(cmds, cmd)
		}
		m.loaded = true
		return m, tea.Batch(cmds...)
	case tea.KeyMsg:
		if m.lists[m.focused].footer != nil && m.lists[m.focused].footer.active {
			result, _ := m.lists[m.focused].Update(msg)
			if _, ok := result.(Column); ok {
				m.lists[m.focused] = result.(Column)
			}
			return m, nil
		}
		switch msg.String() {
		case "!":
			{
				projects.nextProject()
				p.RestoreTerminal()
			}
		case "@":
			{
				projects.previousProject()
				p.RestoreTerminal()
			}
		case "l":
			m.focusRight()
			return m, nil
		case "h":
			m.focusLeft()
			return m, nil
		case "L":
			if m.currColumn().list.Cursor() == int(did) {
				break
			}
			task := m.currColumn().list.SelectedItem().(Task)
			m.lists[m.focused].list.RemoveItem(m.lists[m.focused].list.Cursor())
			DeleteDataInJson(task)
			m.focusRight()
			task.status = m.focused
			AddDataToJson(task)
			m.lists[m.focused].list.InsertItem(len(m.lists[m.focused].list.Items()), task)
		case "H":
			if m.currColumn().list.Cursor() == int(todo) {
				break
			}
			task := m.currColumn().list.SelectedItem().(Task)
			m.lists[m.focused].list.RemoveItem(m.lists[m.focused].list.Cursor())
			DeleteDataInJson(task)
			m.focusLeft()
			task.status = m.focused
			AddDataToJson(task)
			m.lists[m.focused].list.InsertItem(len(m.lists[m.focused].list.Items()), task)
		}
	}
	var cmd tea.Cmd
	result, cmd := m.lists[m.focused].Update(msg)
	if _, ok := result.(Column); ok {
		m.lists[m.focused] = result.(Column)
	}
	return m, cmd
}

func (m *ListModel) View() string {
	if m.loaded {
		listModel := lipgloss.JoinHorizontal(lipgloss.Center,
			m.lists[todo].View(),
			m.lists[doing].View(),
			m.lists[did].View(),
		)
		return lipgloss.JoinVertical(lipgloss.Center, m.header, listModel, m.lists[m.focused].foot)
	} else {
		return "hold on now"
	}
}

func (m *ListModel) initializeLists() []Column {
	return renderColumns(GetDataFromJson())
}

func (m *ListModel) focusRight() {
	m.lists[m.focused].focused = false
	m.focused = (m.focused + 1) % 3
	m.lists[m.focused].focused = true
}

func (m *ListModel) focusLeft() {
	if m.focused == todo {
		m.lists[m.focused].focused = false
		m.focused = did
		m.lists[m.focused].focused = true
	} else {
		m.lists[m.focused].focused = false
		m.focused = (m.focused - 1) % 3
		m.lists[m.focused].focused = true
	}
}

func (m *ListModel) currColumn() Column {
	return m.lists[m.focused]
}

func initializeModel() *ListModel {
	return &ListModel{}
}
