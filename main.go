package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("\n***** Welcome to TODO app *****")

	command := flag.String("command", "No command", "Command to run")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	if *command == "create-task" {
		var task, category, duedate string

		print("\nPlease Enter Task: ")
		scanner.Scan()
		task = scanner.Text()

		print("\nPlease Enter Category: ")
		scanner.Scan()
		category = scanner.Text()

		print("\nPlease Enter DueDate: ")
		scanner.Scan()
		duedate = scanner.Text()

		fmt.Printf("\nTask Created Successfuly: %s: %s On %s", category, task, duedate)
	}
	if *command == "create-category" {
		var title, color string

		print("\nPlease Enter Category Title: ")
		scanner.Scan()
		title = scanner.Text()

		print("\nPlease Enter Category Color: ")
		scanner.Scan()
		color = scanner.Text()

		fmt.Printf("\nCategory Created Successfuly: %s | %s", title, color)
	}
}