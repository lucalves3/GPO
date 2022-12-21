package gpo

import (
	"fmt"
	"os"
)

func (c *CreateStructure) CreateMainFileExample(projectName string) {
	path, _ := os.Getwd()
	pathToFile := fmt.Sprintf("%s/%s", path, "/")
	nameFileWithPath := fmt.Sprintf("%s/%s.go", pathToFile, "main")
	_, err := os.Stat(nameFileWithPath)
	if err != nil {
		file, _ := os.Create(nameFileWithPath)
		defer file.Close()
		fileText := fmt.Sprintf(`package main

		import (
			"%s/models"
			"%s/routes"
			"fmt"
		)
		
		var tokenDefense string
		
		func main() {
		
			err := models.CreateConnectionDB()
			if err != nil {
				fmt.Println("[ERROR]: Error to create connection to DB")
			}
		
			routes.HandleRequests()
		
		}
		`, projectName, projectName)

		_, err = file.WriteString(fileText)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = file.Sync()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
