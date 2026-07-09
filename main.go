package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID int
	Name string
	Email string
	Password string
}

type Task struct {
	ID int
	Title string
	DueDate string
	CategoryID int
	IsDone bool
	UserID int

}

type Category struct {
	ID int
	Title string
	Color string
	UserID int
}

var userStorage []User
var taskList []Task
var categoryList []Category
var AuthenticatedUser *User

const userStoragePath = "user.txt"

func main() {

	loadUserStorageFromFile()

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

	if command != "register" && command != "exit" && AuthenticatedUser == nil {
		LoginUser()

		if AuthenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		CreateTask()
	case "create-category":
		CreateCategory()
	case "register":
		RegisterUser()
	case "list-task":
		ListTask()
	case "login":
		if AuthenticatedUser != nil {
			return
		}
		LoginUser()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command Is Not Valid.")
	}
}

func CreateTask() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, category, duedate string

	print("\nPlease Enter Task: ")
	scanner.Scan()
	title = scanner.Text()

	print("\nPlease Enter Category: ")
	scanner.Scan()
	category = scanner.Text()

	CategoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("Category input is not valid, %v\n", err)

		return
	}

	isFound := false
	for _, c := range categoryList {
		if c.ID == CategoryID && c.UserID == AuthenticatedUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Println("Category not found.")

		return
	}

	print("\nPlease Enter DueDate: ")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task {
		ID: len(taskList) + 1,
		Title: title,
		DueDate: duedate,
		CategoryID: CategoryID,
		IsDone: false,
		UserID: AuthenticatedUser.ID,
	}

	taskList = append(taskList, task)

	fmt.Printf("\nTask Created Successfuly: %s | %s On %s", category, title, duedate)
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

	c := Category{
		ID: len(categoryList) + 1,
		Title: title,
		Color: color,
		UserID: AuthenticatedUser.ID,
	}

	categoryList = append(categoryList, c)

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

	var file *os.File

	file, err := os.OpenFile(userStoragePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error while creationg or opening the user.txt file. %s", err)

		return
	}

	data := fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", 
	user.ID, user.Name, user.Email, user.Password)

	var b = []byte(data)

	file.Write(b)

	file.Close()

	fmt.Print("\nUser Created Successfuly")
}

func LoginUser() {
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	print("\nPlease Enter Your Email: ")
	scanner.Scan()
	email = scanner.Text()

	print("\nPlease Enter Your password: ")
	scanner.Scan()
	password = scanner.Text()
	// get the email and password from the client

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			AuthenticatedUser = &user

			break
		}
	}

	if AuthenticatedUser == nil {
		fmt.Println("The Email Or Password Is Not Currect")
	}
}


func ListTask() {
	for _, task := range taskList {
		if task.UserID == AuthenticatedUser.ID {
			fmt.Printf("%+v\n", task)
		}
	}
}

func loadUserStorageFromFile() {
	file, err := os.Open(userStoragePath)
	if err != nil {
		fmt.Printf("Error occurred while opening uset.txt file. %s\n", err)
	}

	var data = make([]byte, 10240)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Printf("Error occurred while reading uset.txt file. %s\n", oErr)
	}

	var dataStr = string(data)
	userSlice := strings.Split(dataStr, "\n")
	for _, u := range userSlice {
		if u == "" {
			continue
		}

		user := User{}
		userFields := strings.Split(u, ",")
		for _, field := range userFields {
			values := strings.Split(field, ": ")
			fieldName := strings.ReplaceAll(values[0], " ", "")
			fieldValue := values[1]

			switch fieldName {
			case "id":
				id, err := strconv.Atoi(fieldValue)
				if err != nil {
					fmt.Println("Strconv error", err)

					return
				}
				user.ID = id
			case "name":
				user.Name = fieldValue
			case "email":
				user.Email = fieldValue
			case "password":
				user.Password = fieldValue
			}
		}
		fmt.Printf("user: %+v\n", user)
	}

}