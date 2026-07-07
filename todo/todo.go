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
	if err == nil {
		return nil
	}
	return json.Unmarshal(data, &t.Tasks)
}
