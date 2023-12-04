package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func AddDataToJson(t Task) {
	jsonData := readDataFromJson(projects.currProject().filePath)

	for project := range jsonData {
		items := jsonData[project]
		for category := range items {
			if category != StatusToString(t.status) {
				continue
			}
			tasks := items[category]
			taskMap := t.toMap()
			tasks = append(tasks, taskMap)
			items[category] = tasks
		}
	}

	marshaledData, _ := json.Marshal(jsonData)
	file, err := os.Create(projects.currProject().filePath)
	if err != nil {
		panic(err)
	}
	file.Write(marshaledData)
	return

}

func EditDataInJson(newTask, oldTask Task) {
	jsonData := readDataFromJson(projects.currProject().filePath)

	for project := range jsonData {
		items := jsonData[project]
		category := StatusToString(oldTask.status)
		for task := range items[category] {
			if oldTask.title == items[category][task]["title"] {
				items[category][task] = newTask.toMap()
			}
		}
	}

	marshaledData, _ := json.Marshal(jsonData)
	file, err := os.Create(projects.currProject().filePath)
	if err != nil {
		panic(err)
	}
	file.Write(marshaledData)
	return
}

func DeleteDataInJson(task Task) {
	jsonData := readDataFromJson(projects.currProject().filePath)

	for projects := range jsonData {
		project := jsonData[projects]
		category := StatusToString(task.status)
		var index int
		for item := range project[category] {
			if task.title == project[category][item]["title"] && task.description == project[category][item]["description"] {
				index = item
			}
		}
		project[category] = removeItem(project[category], index)
	}

	marshaledData, _ := json.Marshal(jsonData)
	file, err := os.Create(projects.currProject().filePath)
	if err != nil {
		panic(err)
	}
	file.Write(marshaledData)
	return

}

func MoveDataInJson(currentIndex, nextIndex int, status string) {
	jsonData := readDataFromJson(projects.currProject().filePath)

	for projects := range jsonData {
		project := jsonData[projects]
		items := project[status]
		swapper := reflect.Swapper(items)
		swapper(currentIndex, nextIndex)
	}
	marshaledData, _ := json.Marshal(jsonData)
	file, err := os.Create(projects.currProject().filePath)
	if err != nil {
		panic(err)
	}
	file.Write(marshaledData)
	return
}

func GetDataFromJson() []todoModel {
	jsonData := readDataFromJson(projects.currProject().filePath)
	var data []todoModel
	for p := range jsonData {
		var project todoModel
		project.project = p
		items := jsonData[p]
		for item := range items {
			var itemList []map[string]string
			tasks := items[item]
			for task := range tasks {
				itemList = append(itemList, tasks[task])
			}
			if item == "todo" {
				project.todo = itemList
			} else if item == "doing" {
				project.doing = itemList
			} else {
				project.did = itemList
			}
		}
		data = append(data, project)
	}
	return data
}

func readDataFromJson(path string) map[string]map[string][]map[string]string {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("rawr uwu")
	}
	var jsonData map[string]map[string][]map[string]string
	err = json.Unmarshal([]byte(file), &jsonData)
	if err != nil {
		fmt.Println("uwu rawrr")
	}
	return jsonData

}

func removeItem(slice []map[string]string, index int) []map[string]string {
	return append(slice[:index], slice[index+1:]...)
}
