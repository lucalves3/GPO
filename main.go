package gpo

import (
	"fmt"
	"os"
	"time"
)

type HandleError struct {
	err error
}

func CreateSRCFolder() {
	err := os.Mkdir("/src", 0755)
	if err != nil {
		fmt.Println("Error to create src")
	}
}

func CreateModelsFolder() {
	err := os.Mkdir("/src/models", 0755)
	if err != nil {
		fmt.Println("Error to create models")
	}
}

func CreateControllersFolder() {
	err := os.Mkdir("/src/controllers", 0755)
	if err != nil {
		fmt.Println("Error to create controllers")
	}
}

func CreateEntitiesFolder() {
	err := os.Mkdir("/src/entities", 0755)
	if err != nil {
		fmt.Println("Error to create entities")
	}
}

func CreateRoutesFolder() {
	err := os.Mkdir("/src/routes", 0755)
	if err != nil {
		fmt.Println("Error to create routes")
	}
}

func (r *HandleError) CreateAllProjectStructure() {
	CreateSRCFolder()
	time.Sleep(1 * time.Second)
	CreateModelsFolder()
	time.Sleep(1 * time.Second)
	CreateControllersFolder()
	time.Sleep(1 * time.Second)
	CreateEntitiesFolder()
	time.Sleep(1 * time.Second)
	CreateRoutesFolder()
	time.Sleep(1 * time.Second)
}
