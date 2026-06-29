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

	RunCommand(*command)
} 

func RunCommand(command string) {
	switch command {
	case "create-task":
		CreateTask()
	case "create-category":
		CreateCategory()
	case "register":
		RegisterUser()
	case "login":
		LoginUser()
	default:
		fmt.Println("Command Is Not Valid.")
	}
}

func CreateTask() {
	scanner := bufio.NewScanner(os.Stdin)
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

func CreateCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	print("\nPlease Enter Category Title: ")
	scanner.Scan()
	title = scanner.Text()

	print("\nPlease Enter Category Color: ")
	scanner.Scan()
	color = scanner.Text()

	fmt.Printf("\nCategory Created Successfuly: %s | %s", title, color)
}

func RegisterUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	print("\nPlease Enter Your Email: ")
	scanner.Scan()
	email = scanner.Text()

	id = email

	print("\nPlease Enter Your password: ")
	scanner.Scan()
	password = scanner.Text()

	fmt.Printf("\nUser Created Successfuly: %s | %s | %s", id, email, password)
}

func LoginUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var id, email, password string

	print("\nPlease Enter Your Email: ")
	scanner.Scan()
	email = scanner.Text()

	id = email

	print("\nPlease Enter Your password: ")
	scanner.Scan()
	password = scanner.Text()

	fmt.Printf("\nUser Created Successfuly: %s | %s | %s", id, email, password)
}