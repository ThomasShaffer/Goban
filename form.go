package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Form struct {
	header      string
	title       textinput.Model
	description textinput.Model
	active      bool
    isEdit      bool
	style       lipgloss.Style
}

func NewForm() *Form {
	form := Form{
		header:      "new form",
		title:       textinput.New(),
		description: textinput.New(),
	}
	form.title.Placeholder = "new task"
	form.title.Focus()
	form.description.Placeholder = "add description..."
	form.active = true
	return &form
}


func EditForm(task Task) *Form {
    form := Form{
        header: "edit form",
        title: textinput.New(),
        description: textinput.New(),
    }
    form.title.SetValue(task.title)//(Value = task.title
    form.title.Focus()
    form.description.SetValue(task.description)//(Value = task.title
    form.active = true
    form.isEdit = true
    return &form
}

func (f *Form) Init() tea.Cmd { return nil }
func (f *Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		f.style = lipgloss.NewStyle().Padding(1, 1).BorderStyle(lipgloss.NormalBorder()).Width(msg.Width - 2).Height(msg.Height/3 - 3).Align(lipgloss.Center)
	case tea.KeyMsg:
		if f.title.Focused() {
			switch msg.String() {
			case "tab":
				f.title.Blur()
				f.description.Focus()
				return f, cmd
			}
			result, cmd := f.title.Update(msg)
			f.title = result
			return f, cmd
		}
		switch msg.String() {
		case "tab":
			f.description.Blur()
			f.title.Focus()
			return f, cmd
		}
		result, cmd := f.description.Update(msg)
		f.description = result
		return f, cmd
	}
	return f, cmd
}
func (f *Form) View() string {
	return f.style.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			f.header,
			f.title.View(),
			f.description.View(),
		),
	)
}