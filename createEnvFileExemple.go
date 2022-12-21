package gpo

import (
	"fmt"
	"os"
)

func (c *CreateStructure) CreateEnvFileExample(DBName string) {
	path, _ := os.Getwd()
	pathToFile := fmt.Sprintf("%s/%s", path, "/")
	nameFileWithPath := fmt.Sprintf("%s/%s", pathToFile, ".env")
	_, err := os.Stat(nameFileWithPath)
	if err != nil {
		file, _ := os.Create(nameFileWithPath)
		defer file.Close()
		fileText := fmt.Sprintf(`
		APIPORT=3001
		APIDEV=3002
		DBUSER=root
		DBPASSWORD=senha123456
		DBIP=localhost
		DBPORT=3306
		DBSCHEMA=%s
		SECRET_KEY=SECRET`, DBName)

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
