package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetDataFromJson() []todoModel {
	file, err := os.ReadFile("./items.json")
	if err != nil {
		fmt.Println("rawr uwu")
	}
	var jsonData map[string]map[string][]map[string]string
	err = json.Unmarshal([]byte(file), &jsonData)
	if err != nil {
		fmt.Println("uwu rawrr")
	}

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
