package main

import (
	"fmt"
	"os"
)

type Projects struct {
	name      string
	filePath  string
	directory []string
	index     int
}

func (p *Projects) nextProject() {
	p.index = (p.index + 1) % len(p.directory)
	p.filePath = p.directory[p.index]
}

func (p *Projects) previousProject() {
	if p.index == 0 {
		p.index = len(p.directory) - 1
		p.filePath = p.directory[p.index]
		p.name = p.filePath
	} else {
		p.index = p.index - 1
		p.filePath = p.directory[p.index]
		p.name = p.filePath
	}
}

func initializeProjects() *Projects {
	var projectsDirectory []string
	dir, err := os.Open("./projects")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	files, err := dir.ReadDir(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		projectsDirectory = append(projectsDirectory, "./projects/"+file.Name())
	}
	return &Projects{
		name:      projectsDirectory[0],
		filePath:  projectsDirectory[0],
		directory: projectsDirectory,
		index:     0,
	}
}
