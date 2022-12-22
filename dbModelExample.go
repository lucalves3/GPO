package gps

import (
	"fmt"
	"os"
)

func (c *CreateStructure) CheckIfIndexDatabaseModelWasCreatedAndCreateIfDontExists(projectName string) {
	path, _ := os.Getwd()
	pathToFile := fmt.Sprintf("%s/%s", path, "src/models")
	nameFileWithPath := fmt.Sprintf("%s/%s.go", pathToFile, "db")
	_, err := os.Stat(nameFileWithPath)
	if err != nil {
		file, _ := os.Create(nameFileWithPath)
		defer file.Close()
		fileText := fmt.Sprintf(`package models

		import (
			"%s/src/entities"
			"fmt"
			"os"
		
			"github.com/joho/godotenv"
			"gorm.io/driver/mysql"
			"gorm.io/gorm"
		)
		
		func ConnectDB() (*gorm.DB, error) {
			err := godotenv.Load(".env")
			if err != nil {
				fmt.Println("[ERROR]: Error get env")
			}
			dsn := fmt.Sprint(os.Getenv("DBUSER"), ":", os.Getenv("DBPASSWORD"), "@tcp(", os.Getenv("DBIP"), ":", os.Getenv("DBPORT"), ")/", os.Getenv("DBSCHEMA"), "?charset=utf8mb4&parseTime=True&loc=Local")
			fmt.Println(dsn)

			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		
			if err != nil {
				fmt.Println("[ERROR]: Error creating database")
			}
			ConfigDB(db)
			return db, err
		}
		
		func ConfigDB(db *gorm.DB) {
		
			db.AutoMigrate(&entities.User{}, &entities.Scheduled{}, &entities.Company{})
		}
		
		func CreateConnectionDB() error {
			var err error
			db, err = ConnectDB()
			return err
		}
		`, projectName)

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
