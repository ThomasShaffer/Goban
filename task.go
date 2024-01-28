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
    slice := min(len(t.title), config.taskTitleCutoff)
    return t.title[:slice] 
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
	taskMap["date"] = t.date
	return taskMap

}
