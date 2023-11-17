package main

import "time"

type Task struct {
	status      status
	title       string
	description string
	date        string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	slice := min(len(t.description), 20)
	return t.description[:slice] + "..."
}

func (t Task) Date() string {
	return t.date
}

func (t Task) toMap() map[string]string {
    taskMap := make(map[string]string)
    taskMap["title"] = t.title
    taskMap["description"] = t.description
    currTime := time.Now()
    taskMap["date"] = currTime.Format("01-01-2006")
    return taskMap
    
}
