package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

var tasks []Task
var nextID int = 1

func listTask() {
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks yet. Use 'add' to create one.")
	} else {
		for _, task := range tasks {
			status := "❌"
			if task.Completed {
				status = "✅"
			}
			fmt.Printf("[%d] %s %s\n", task.ID, status, task.Title)
		}
	}
}

func main() {

	err := loadTasks()
	if err != nil {
		fmt.Printf("❌ Failed to load tasks: %v \n", err)
		return
	}

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: todo-cli [add|list|done|delete] [arguments]")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("Please provide task description.")
			return
		}

		title := strings.Join(args[2:], " ")
		newTask := Task{
			ID:        nextID,
			Title:     title,
			Completed: false,
		}
		tasks = append(tasks, newTask)

		nextID++
		saveTasks()
		fmt.Printf("✅ Added task: [%d] %s\n", newTask.ID, newTask.Title)

	case "list":
		listTask()
	case "done":
		markTaskDone(args)
		saveTasks()
	case "delete":
		deleteTask(args)
		saveTasks()
	default:
		fmt.Println("Unknown command: ", command)
	}

}

func markTaskDone(args []string) {
	if len(args) < 3 {
		fmt.Println("Please provide the task ID to mark as done.")
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			fmt.Printf("✅ Task [%d] marked as done.\n", id)
			return
		}

	}
	fmt.Println("Task ID not found.")
}

func deleteTask(args []string) {
	if len(args) < 3 {
		fmt.Println("Please provide the task ID to delete.")
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid task ID.")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("🗑️  Deleted task [%d]: %s\n", id, task.Title)
			return
		}

	}
	fmt.Println("Task ID not found.")
}

func saveTasks() error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func loadTasks() error {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	nextID = maxID + 1

	return nil
}
