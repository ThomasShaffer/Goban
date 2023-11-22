package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadDataFromJson(path string) map[string]map[string][]map[string]string {
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

func WriteDataToJson(t Task) {
	jsonData := ReadDataFromJson("./items.json")

	for project := range jsonData {
		items := jsonData[project]
		for category := range items {
			if category != "todo" {
				continue
			}
			tasks := items[category]
			taskMap := t.toMap()
			tasks = append(tasks, taskMap)
			items[category] = tasks
		}
	}

	marshaledData, _ := json.Marshal(jsonData)
	file, err := os.Create("./items.json")
	if err != nil {
		panic(err)
	}
	file.Write(marshaledData)
	return

}

func EditDataInJson(newTask, oldTask Task) {
	jsonData := ReadDataFromJson("./items.json")

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
	file, err := os.Create("./items.json")
	if err != nil {
		panic(err)
	}
	file.Write(marshaledData)
	return
}

func DeleteDataInJson(task Task) {
	jsonData := ReadDataFromJson("./items.json")

	for project := range jsonData {
		items := jsonData[project]
		category := StatusToString(task.status)
		for item := range items[category] {
			if task.title == items[category][item]["title"] &&
				task.description == items[category][item]["description"] {
				items[category] = removeItem(items[category], item)
			}
		}
	}

	marshaledData, _ := json.Marshal(jsonData)
	file, err := os.Create("./items.json")
	if err != nil {
		panic(err)
	}
	file.Write(marshaledData)
	return

}

func GetDataFromJson() []todoModel {
	jsonData := ReadDataFromJson("./items.json")
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

func removeItem(slice []map[string]string, index int) []map[string]string {
	return append(slice[:index], slice[index+1:]...)
}
