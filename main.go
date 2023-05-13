package main

import (
	"fmt"
	"os"
	"todoapp/controllers"

	"github.com/gin-gonic/gin" // Import Redis package
	"github.com/joho/godotenv"

	// docs "todoapp/docs"
	seed "todoapp/seeder"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger todoapp APP
// @version         2.0
// @description     This is a swagger documentation for Costumer APP.
// @BasePath        /api/v1
// @host            localhost:8080
// @schemes         http https
// @SecurityDefinition  jwt
// @Security        jwt
func main() {

	//setting env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	//migrate and seeder
	seed.CreateMigration()
	seed.SeedActivities()

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// docs.SwaggerInfo.BasePath = "/api/v1"

	activitiesController := controllers.ActivitiesController{}

	v1 := router.Group("")

	v1.GET("/activity-groups", activitiesController.GetActivities)

	v1.POST("/activity-groups", activitiesController.CreateActivity)
	v1.PUT("/activity-groups/:id", activitiesController.UpdateActivity)
	v1.DELETE("/activity-groups/:id", activitiesController.DeleteActivity)
	v1.GET("/activity-groups/:id", activitiesController.GetActivity)

	router.Run(":3030")
}
