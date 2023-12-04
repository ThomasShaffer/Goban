package main

import (
	"fmt"
	"os"
)

type Projects struct {
	directory []File
	index     int
}

type File struct {
	name     string
	filePath string
}

func (p *Projects) nextProject() {
	p.index = (p.index + 1) % len(p.directory)
}

func (p *Projects) previousProject() {
	if p.index == 0 {
		p.index = len(p.directory) - 1
	} else {
		p.index = p.index - 1
	}
}

func (p *Projects) currProject() File {
	return p.directory[p.index]
}

func initializeProjects() *Projects {
	var projects []File
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
		projects = append(projects,
			File{
				name:     file.Name(),
				filePath: "./projects/" + file.Name(),
			})
	}
	return &Projects{
		directory: projects,
		index:     0,
	}
}
