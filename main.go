package main

import (
	"fmt"
	"os"
	"strconv"
	"todo-app/todo"
)

func main() {
	todoList := todo.ToDoList{FilePath: "todo.json"}
	todoList.Load()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command: add, remove, list")
		return
	}
	command := os.Args[1]
	switch command {
	case "add":
		todoList.Add(os.Args[2])
	case "remove":
		s, _ := strconv.Atoi(os.Args[2])
		todoList.Remove(s)
	case "list":
		todoList.PrintList()
	default:
		fmt.Println("Unknown command")
	}
}
