package todo

import (
	"encoding/json"
	"os"
)

type Task struct {
	Id    int
	Title string
}

type ToDoList struct {
	Tasks    []Task
	FilePath string
}

func (t *ToDoList) Load() error {
	data, err := os.ReadFile(t.FilePath)
	if err != nil {
		return nil
	}
	return json.Unmarshal(data, &t.Tasks)
}

func (t *ToDoList) Save() error {
	data, err := json.Marshal(t.Tasks)
	if err == nil {
		err = os.WriteFile(t.FilePath, data, 0644)
	}
	return err
}
