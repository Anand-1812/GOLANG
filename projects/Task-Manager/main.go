package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"Task-Manager/addtask"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide the options for task :)")
	}

	fmt.Println("---- CLI Task Manager ----")
	command := os.Args[1]

	switch command {
	case "Add":
		task := strings.Join(os.Args[2:], " ")
		f, err := os.OpenFile("task.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
		if err := addtask.Add(f, task);err != nil {
			log.Fatal(err)
		}
	}
}
