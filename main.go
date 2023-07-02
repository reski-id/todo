package main

import (
	"todoapp/controllers"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	// Create a new Graylog hook
	hook := graylog.NewGraylogHook("localhost:12201", nil)

	// Set the formatter to JSON
	log.Formatter = new(logrus.JSONFormatter)

	// Add the Graylog hook to the logger
	log.Hooks.Add(hook)

	// Set the log level
	log.SetLevel(logrus.DebugLevel)

	router := gin.Default()

	activitiesController := controllers.NewActivitiesController(log)

	v1 := router.Group("")

	v1.GET("/activity-groups", activitiesController.GetActivities)
	v1.POST("/activity-groups", activitiesController.CreateActivity)
	v1.PUT("/activity-groups/:id", activitiesController.UpdateActivity)
	v1.DELETE("/activity-groups/:id", activitiesController.DeleteActivity)
	v1.GET("/activity-groups/:id", activitiesController.GetActivity)

	router.Run(":3030")
}
