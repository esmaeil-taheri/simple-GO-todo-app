package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID int
	Name string
	Email string
	Password string
}

var userStorage []User

func main() {
	fmt.Println("\n***** Welcome to TODO app *****")

	command := flag.String("command", "No command", "Command to run")
	flag.Parse()

	for {
		RunCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\nPlease Enter Another Command: ")
		scanner.Scan()
		*command = scanner.Text()
	}

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
	case "exit":
		os.Exit(0)
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
	var name, email, password string

	print("\nPlease Enter Your Email: ")
	scanner.Scan()
	email = scanner.Text()

	print("\nPlease Enter Your Name: ")
	scanner.Scan()
	name = scanner.Text()

	print("\nPlease Enter Your password: ")
	scanner.Scan()
	password = scanner.Text()

	user := User {
		ID: len(userStorage) + 1,
		Name: name,
		Email: email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	fmt.Print("\nUser Created Successfuly")
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