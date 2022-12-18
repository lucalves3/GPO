package gpo

import (
	"os"
	"time"
)

type HandleError struct {
	err error
}

func (r *HandleError) createSRCFolder() error {
	err := os.Mkdir("/src", 0755)
	return err
}

func (r *HandleError) createModelsFolder() error {
	err := os.Mkdir("/src/models", 0755)
	return err
}

func (r *HandleError) createControllersFolder() error {
	err := os.Mkdir("/src/controllers", 0755)
	return err
}

func (r *HandleError) createEntitiesFolder() error {
	err := os.Mkdir("/src/entities", 0755)
	return err
}

func (r *HandleError) createRoutesFolder() error {
	err := os.Mkdir("/src/routes", 0755)
	return err
}

func (r *HandleError) CreateAllProjectStructure() error {
	r.createSRCFolder()
	time.Sleep(1 * time.Second)
	r.createModelsFolder()
	time.Sleep(1 * time.Second)
	r.createControllersFolder()
	time.Sleep(1 * time.Second)
	r.createEntitiesFolder()
	time.Sleep(1 * time.Second)
	err := r.createRoutesFolder()
	time.Sleep(1 * time.Second)
	return err
}
