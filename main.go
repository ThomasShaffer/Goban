package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

var m *ListModel
var projects *Projects
var p *tea.Program

func main() {
	projects = initializeProjects()
	m = initializeModel()
	p = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
