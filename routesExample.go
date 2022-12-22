package gps

import (
	"fmt"
	"os"
)

func (c *CreateStructure) CreateRoutesFileExample(projectName string) {
	path, _ := os.Getwd()
	pathToFile := fmt.Sprintf("%s/%s", path, "src/routes")
	nameFileWithPath := fmt.Sprintf("%s/%s", pathToFile, "routes.go")
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
				%s := main.Group("YOUR_ROUTE_HERE")
				{
					users.GET("/", controllers.GetAll%s)
					users.GET("/:id", controllers.Get%sByID)
					users.POST("/", controllers.Create%s)
					users.DELETE("/:id", controllers.Delete%s)
				}
			router.Run(fmt.Sprint(":", os.Getenv("APIPORT")))
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
