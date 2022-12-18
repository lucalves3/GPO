package gpo

import (
	"os"
	"time"
)

type HandleError struct {
	err error
}

func createSRCFolder() error {
	err := os.Mkdir("/src", 0755)
	return err
}

func createModelsFolder() error {
	err := os.Mkdir("/src/models", 0755)
	return err
}

func createControllersFolder() error {
	err := os.Mkdir("/src/controllers", 0755)
	return err
}

func createEntitiesFolder() error {
	err := os.Mkdir("/src/entities", 0755)
	return err
}

func createRoutesFolder() error {
	err := os.Mkdir("/src/routes", 0755)
	return err
}

func CreateAllProjectStructure() error {
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
