package main

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
