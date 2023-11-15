package main

import "github.com/charmbracelet/bubbles/list"

type status int

const (
	todo status = iota
	doing
	did
)

type todoModel struct {
	project string
	todo    []map[string]string
	doing   []map[string]string
	did     []map[string]string
}

type ListModel struct {
	focused status
	lists   []list.Model
	loaded  bool
}

type Task struct {
	status      status
	title       string
	description string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}
