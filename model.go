package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListModel struct {
	focused status
	lists   []Column
	loaded  bool
}

func (m *ListModel) Init() tea.Cmd { return nil }
func (m *ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		var cmds []tea.Cmd
		titles := []string{"todo", "doing", "did"}
		m.initializeLists(titles)
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
		switch msg.String() {
		case "l":
			m.lists[m.focused].focused = false
			m.focused = (m.focused + 1) % 3
			m.lists[m.focused].focused = true
			return m, nil
		case "h":
			if m.focused == todo {
				m.lists[m.focused].focused = false
				m.focused = did
				m.lists[m.focused].focused = true
				return m, nil
			} else {
				m.lists[m.focused].focused = false
				m.focused = (m.focused - 1) % 3
				m.lists[m.focused].focused = true
			}
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
		m.lists[description].list.SetSize(0, 15)
		return lipgloss.JoinVertical(lipgloss.Center, "\nGoban", listModel, m.lists[description].View())
	} else {
		return "hold on now"
	}
}

func (m *ListModel) initializeLists(title []string) {
	m.lists = []Column{NewColumn(title[0], 0, 0), NewColumn(title[1], 0, 0), NewColumn(title[2], 0, 0), NewBottomColumn(0, 0)}
}

func initializeModel() *ListModel {
	return &ListModel{}
}
