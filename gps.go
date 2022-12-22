package gps

import (
	"fmt"
	"os"
	"time"
)

type CreateStructure struct {
	gps *CreateStructure
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
