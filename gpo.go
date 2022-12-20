package gpo

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type CreateStructure struct {
	err error
}

func createSRCFolder() error {
	path, _ := os.Getwd()
	localFile := fmt.Sprintf("%s/%s", path, "src")
	err := os.Mkdir(localFile, 0755)
	return err
}

func createModelsFolder() error {
	path, _ := os.Getwd()
	localFile := fmt.Sprintf("%s/%s", path, "src/models")
	err := os.Mkdir(localFile, 0755)
	return err
}

func createControllersFolder() error {
	path, _ := os.Getwd()
	localFile := fmt.Sprintf("%s/%s", path, "src/controllers")
	err := os.Mkdir(localFile, 0755)
	return err
}

func createEntitiesFolder() error {
	path, _ := os.Getwd()
	localFile := fmt.Sprintf("%s/%s", path, "src/entities")
	err := os.Mkdir(localFile, 0755)
	return err
}

func createRoutesFolder() error {
	path, _ := os.Getwd()
	localFile := fmt.Sprintf("%s/%s", path, "src/routes")
	err := os.Mkdir(localFile, 0755)
	return err
}

func (c *CreateStructure) CreateAllProjectStructureFolders() error {
	createSRCFolder()
	time.Sleep(1 * time.Second)
	createModelsFolder()
	time.Sleep(1 * time.Second)
	createControllersFolder()
	time.Sleep(1 * time.Second)
	createEntitiesFolder()
	time.Sleep(1 * time.Second)
	err := createRoutesFolder()
	time.Sleep(1 * time.Second)
	return err
}

func (c *CreateStructure) CreateModelExampleFile(projectName string, modelName string) {
	path, _ := os.Getwd()
	localFile := fmt.Sprintf("%s/%s", path, "src/models")
	firstCapitalCase := strings.Title(modelName)
	modelNameString := fmt.Sprintf("%s/%s.go", localFile, modelName)
	file, err := os.Create(modelNameString)
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
