package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

func (m *ListModel) Init() tea.Cmd { return nil }
func (m *ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initializeLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "l":
			m.focused = (m.focused + 1) % 3
			return m, nil
		case "h":
			if m.focused == todo {
				m.focused = did
				return m, nil
			} else {
				m.focused = (m.focused - 1) % 3
			}
		}
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m *ListModel) View() string {
	if m.loaded {
		maxHeight := lipgloss.NewStyle().GetMaxHeight()
		maxWidth := lipgloss.NewStyle().GetMaxWidth()
		return lipgloss.JoinHorizontal(lipgloss.Center,
			lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Width(maxWidth/3).Height(maxHeight-1).Render(m.lists[todo].View()),
			lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Width(maxWidth/3).Height(maxHeight-1).Render(m.lists[doing].View()),
			lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).Width(maxWidth/3).Height(maxHeight-1).Render(m.lists[did].View()),
		)
	} else {
		return "hold on now"
	}
}

func (m *ListModel) initializeLists(width, height int) {
	didDelegate := list.NewDefaultDelegate()
	didDelegate.ShowDescription = false

	otherDelegate := list.NewDefaultDelegate()

	didItems := list.New([]list.Item{}, didDelegate, width, height)
	otherItems := list.New([]list.Item{}, otherDelegate, width, height)
	didItems.SetShowHelp(false)
	otherItems.SetShowHelp(false)

	m.lists = []list.Model{otherItems, otherItems, didItems}

	m.lists[todo].Title = "todo"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "do this", description: "now"},
		Task{status: todo, title: "do this 1", description: "ok"},
		Task{status: todo, title: "do this 2", description: "yerr"},
	})

	m.lists[doing].Title = "doing"
	m.lists[doing].SetItems([]list.Item{
		Task{status: doing, title: "doing this", description: "now"},
		Task{status: doing, title: "doing this 1", description: "nw"},
		Task{status: doing, title: "doing this 2", description: "nsdow"},
	})

	m.lists[did].Title = "did"
	m.lists[did].SetItems([]list.Item{
		Task{status: did, title: "did this", description: "now"},
		Task{status: did, title: "did this 1", description: "npok"},
		Task{status: did, title: "did this 2", description: "into the darkness we search"},
	})

}

func initializeModel() *ListModel {
	return &ListModel{}
}

func main() {
	m := initializeModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
