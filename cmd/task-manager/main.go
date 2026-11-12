
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Task struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsPending   bool   `json:"ispending"`
	Time        string `json:"time"`
}

func getJsonInfo() []Task {
	var tasks []Task

	file, err := os.Open("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			return tasks
		}
		fmt.Println("Error:", err)
		return tasks
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		if err.Error() == "EOF" {
			return tasks
		}
		fmt.Println("Error decoding JSON:", err)
		return tasks
	}

	return tasks
}

func saveTasks(tasks []Task) error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	return encoder.Encode(tasks)
}

func createTask(name, description string) *Task {
	existingTasks := getJsonInfo()

	var maxID int64

	for _, t := range existingTasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	task := Task{
		ID:          maxID + 1,
		Name:        name,
		Description: description,
		IsPending:   true,
		Time:        time.Now().Format("2006-01-02 15:04:05"),
	}

	existingTasks = append(existingTasks, task)

	if err := saveTasks(existingTasks); err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	fmt.Println("Created:", task.Name)

	return &task
}

func addTask(args []string) {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)

	name := fs.String("name", "", "The task name")
	description := fs.String("description", "", "The description of the task")

	if err := fs.Parse(args); err != nil {
		return
	}

	if strings.TrimSpace(*name) == "" {
		fmt.Println("The name is required.")
		fmt.Println("Usage: task-manager add -name <name>")
		return
	}

	createTask(*name, *description)
}


func taskList() {
    tasks := getJsonInfo()

    if len(tasks) == 0 {
        fmt.Println("No tasks found.")
        return
    }

    fmt.Println("+----+----------------+-----------+---------------------+")
    fmt.Printf("| %-2s | %-14s | %-9s | %-19s |\n",
        "ID", "TITLE", "STATUS", "CREATED AT")
    fmt.Println("+----+----------------+-----------+---------------------+")

    for _, task := range tasks {
        status := "pending"

        if !task.IsPending {
            status = "completed"
        }

        fmt.Printf("| %-2d | %-14s | %-9s | %-19s |\n",
            task.ID,
            task.Name,
            status,
            task.Time,
        )
    }

    fmt.Println("+----+----------------+-----------+---------------------+")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-manager <command>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		addTask(os.Args[2:])
	case "list":
		taskList()
	default:
		fmt.Println("Unknown command:", command)
	}
}
