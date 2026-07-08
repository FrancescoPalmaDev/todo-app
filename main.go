package main

import (
	"fmt"
	"os"
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
		fmt.Println("This task have been added to the list")
	case "remove":
		fmt.Println("This task have been removed from the list")
	case "list":
		fmt.Println("Here is the list of tasks")
	default:
		fmt.Println("Unknown command")
	}
}
