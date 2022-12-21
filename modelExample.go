package gpo

import (
	"fmt"
	"os"
	"strings"
)

func (c *CreateStructure) CreateModelExampleFile(projectName string, modelName string) {
	path, _ := os.Getwd()
	localFile := fmt.Sprintf("%s/%s", path, "src/models")
	firstCapitalCase := strings.Title(modelName)
	modelNameString := fmt.Sprintf("%s/%s.go", localFile, modelName)
	file, err := os.Create(modelNameString)
	CheckIfIndexModelWasCreated(projectName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileText := fmt.Sprintf(`package models

	import (
		"%s/entities"
	
		"gorm.io/gorm"
	)
	
	var db *gorm.DB
	
	func Create%s(%s entities.%s) error {
		tx := db.Begin()
		result := tx.Create(&%s)
		if result.Error != nil {
			return result.Error
		}
		tx.Commit()
		return nil
	}
	
	func Get%sByEmail(email any) (entities.%s, error) {
		var %s entities.%s
		result := db.Where("email = ? ", email).First(&%s)
		return %s, result.Error
	}
	
	func Get%sById(Id uint64) (entities.%s, error) {
		var %s entities.%s
		result := db.Where("id = ? ", Id).First(&%s)
		return %s, result.Error
	}
	
	func GetAll%ss() ([]entities.%s, error) {
		var %ss []entities.%s
		result := db.Model(&%ss).Find(&%ss)
		return %ss, result.Error
	}
	
	func Delete%s(id uint64) error {
		%s, err := Get%sById(id)
		if err != nil {
			return err
		}
		tx := db.Begin()
		err = tx.Model(&%s).Replace(nil)
		if err != nil {
			return err
		}
		result := tx.Delete(&%s)
		if result.Error != nil {
			return result.Error
		}
		tx.Commit()
		return nil
	}`, projectName, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase, firstCapitalCase)

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

func CheckIfIndexModelWasCreated(projectName string) {
	path, _ := os.Getwd()
	pathToFile := fmt.Sprintf("%s/%s", path, "src/models")
	nameFileWithPath := fmt.Sprintf("%s/%s.go", pathToFile, "index")
	_, err := os.Stat(nameFileWithPath)
	if err != nil {
		file, _ := os.Create(nameFileWithPath)
		defer file.Close()
		fileText := fmt.Sprintf(`package models

		import (
			"gorm.io/gorm"
		)
		
		var db *gorm.DB`)

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
