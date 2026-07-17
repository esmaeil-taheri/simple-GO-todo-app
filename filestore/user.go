package filestore

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"errors"
	"todoapp/entity"
	"todoapp/constant"
	"strconv"
)


type FileStore struct {
	filePath string
	serializationMode string
}

// constractor
func New(path, serializationMode string) FileStore {
	return FileStore{filePath: path, serializationMode: serializationMode}
}

func (f FileStore) Save(u entity.User) {
	f.writeUserToFile(u)
}


func (f FileStore) Load() []entity.User {
	var uStore []entity.User

	file, err := os.Open(f.filePath)
	if err != nil {
		fmt.Printf("Error occurred while opening uset.txt file. %s\n", err)
	}

	var data = make([]byte, 10240)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Printf("Error occurred while reading uset.txt file. %s\n", oErr)

		return nil
	}

	var dataStr = string(data)

	userSlice := strings.Split(dataStr, "\n")
	for _, u := range userSlice {
		var userStruct = entity.User{}

		switch f.serializationMode {
		case constant.TextMode:
			var dErr error
			userStruct, dErr = deserializeFromText(u)
			if dErr != nil {
				fmt.Println("Cannot deserialize user record to user struct in text mode.", dErr)
				
				return nil
			}
		case constant.JsonMode:
			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("Cannot deserialize user record to user struct in json mode.", uErr)
				
				return nil
			}
		default:
			fmt.Println("invalid serialization mode")

			return nil
		}
		
		uStore = append(uStore, userStruct)
	}

	return uStore
}

func (f FileStore) writeUserToFile(user entity.User) {

	var file *os.File

	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error while creationg or opening the user.txt file. %s", err)

		return
	}
	defer file.Close()

	var data []byte
	if f.serializationMode == constant.TextMode {
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n", 
			user.ID, user.Name, user.Email, user.Password))
	} else if f.serializationMode == constant.JsonMode {
		var jErr error
		data, jErr = json.Marshal(user)
		if jErr != nil {
			fmt.Println("Error occurred while marshaling user to json.", jErr)

			return
		}
	} else {
		fmt.Println("Invalid serialization mode.")

		return
	}

	data = append(data, []byte("\n")...)

	file.Write(data)

	fmt.Print("\nUser Created Successfuly")
}

func deserializeFromText(userStr string) (entity.User, error) {
	if userStr == "" {

			return entity.User{}, errors.New("user string is empty.")
		}

		var user = entity.User{}
		userFields := strings.Split(userStr, ",")
		for _, field := range userFields {
			values := strings.Split(field, ": ")
			if len(values) != 2{
				fmt.Println("field is not valid, skipping...", len(values))

				continue
			}
			fieldName := strings.ReplaceAll(values[0], " ", "")
			fieldValue := values[1]

			switch fieldName {
			case "id":
				id, err := strconv.Atoi(fieldValue)
				if err != nil {
					return entity.User{}, errors.New("Str conv error.")
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

		return user, nil
}