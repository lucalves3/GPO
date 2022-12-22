package gps

import (
	"fmt"
	"os"
	"strings"
)

func (c *CreateStructure) CreateRoutesFileExample(projectName string, entitieName string) {
	path, _ := os.Getwd()
	pathToFile := fmt.Sprintf("%s/%s", path, "src/routes")
	nameFileWithPath := fmt.Sprintf("%s/%s", pathToFile, "routes.go")
	firstCapitalCase := strings.Title(entitieName)
	_, err := os.Stat(nameFileWithPath)
	if err != nil {
		file, _ := os.Create(nameFileWithPath)
		defer file.Close()
		fileText := fmt.Sprintf(`package routes

		import (
			"%s/src/controllers"
			"fmt"
			"os"
		
			"github.com/gin-contrib/cors"
			"github.com/gin-gonic/gin"
			"github.com/joho/godotenv"
		)
		
		func HandleRequests() {
			err := godotenv.Load(".env")
			if err != nil {
				fmt.Println("[ERROR]: Error get env")
			}
		
			router := gin.Default()
		
			config := cors.DefaultConfig()
			config.AllowOrigins = []string{"*"}
		
			router.Use(cors.New(config))
		
			main := router.Group("api/")
			{
				%s := main.Group("%s")
				{
					%s.GET("/", controllers.GetAll%s)
					%s.GET("/:id", controllers.Get%sByID)
					%s.POST("/", controllers.Create%s)
					%s.DELETE("/:id", controllers.Delete%s)
				}
			router.Run(fmt.Sprint(":", os.Getenv("APIPORT")))
		}
	}
		`, projectName, entitieName, entitieName, entitieName, firstCapitalCase, entitieName, firstCapitalCase, entitieName, firstCapitalCase, entitieName, firstCapitalCase)

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
