package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"Task-Manager/addtask"
	"Task-Manager/deleteTask"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide the options for task :)")
		PrintUsage()
		os.Exit(1)
	}

	fmt.Println("---- CLI Task Manager ----")
	command := os.Args[1]

	switch strings.ToLower(command) {
	case "add":
		if len(os.Args) < 3 {
			log.Fatal("provide a task to add")
		}

		task := strings.Join(os.Args[2:], " ")
		f, err := os.OpenFile("task.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
		if err := addtask.Add(f, task); err != nil {
			log.Fatal(err)
		}

	case "delete":
		if len(os.Args) < 3 {
			log.Fatal("provide the task number to delete")
		}

		if err := deletetask.Delete("task.txt", os.Args[2]); err != nil {
			log.Fatal(err)
		}
	case "print", "list":
		if err := PrintTasks("task.txt"); err != nil {
			log.Fatal(err)
		}
	default:
		PrintUsage()
	}
}

func PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("   add <task>       - Add a task in the list")
	fmt.Println("   delete <number>  - Delete a task from the list")
	fmt.Println("   print            - Print all tasks")
}

func PrintTasks(fileName string) error {
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		fmt.Println("No tasks found")
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	taskNumber := 1
	for scanner.Scan() {
		task := strings.TrimSpace(scanner.Text())
		if task == "" {
			continue
		}

		fmt.Printf("%d. %s\n", taskNumber, task)
		taskNumber++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	if taskNumber == 1 {
		fmt.Println("No tasks found")
	}

	return nil
}
