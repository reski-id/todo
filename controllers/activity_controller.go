package controllers

import (
	"net/http"
	"strconv"
	"time"
	"todoapp/models"
	"todoapp/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ActivitiesController struct {
	logger *logrus.Logger
}

func NewActivitiesController(logger *logrus.Logger) *ActivitiesController {
	return &ActivitiesController{logger: logger}
}

func (controller *ActivitiesController) CreateActivity(c *gin.Context) {
	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	var activity models.Activitie
	if err := c.ShouldBind(&activity); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if result := db.Create(&activity); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}
	response := models.ActivitieResponseSingle{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	}
	c.JSON(http.StatusCreated, response)
}

func (controller *ActivitiesController) GetActivities(c *gin.Context) {
	// all users can access
	db, err := utils.Connect()
	if err != nil {
		errorMessage := "Failed to connect to the database"
		controller.logger.WithFields(logrus.Fields{
			"timestamp": time.Now().Format(time.RFC3339),
			"level":     "error",
			"message":   errorMessage,
			"source":    "ActivitiesController",
			"tag":       "GetActivities",
			"field":     err.Error(),
		}).Error(errorMessage)

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	var activities []models.Activitie
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset := (page - 1) * limit

	if result := db.Offset(offset).Limit(limit).Find(&activities); result.Error != nil {
		errorMessage := "Failed to fetch activities from the database"
		controller.logger.WithFields(logrus.Fields{
			"timestamp": time.Now().Format(time.RFC3339),
			"level":     "error",
			"message":   errorMessage,
			"source":    "ActivitiesController",
			"tag":       "GetActivities",
			"field":     result.Error.Error(),
		}).Error(errorMessage)

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}

	var total int64
	if result := db.Model(&models.Activitie{}).Count(&total); result.Error != nil {
		errorMessage := "Failed to get the total count of activities from the database"
		controller.logger.WithFields(logrus.Fields{
			"timestamp": time.Now().Format(time.RFC3339),
			"level":     "error",
			"message":   errorMessage,
			"source":    "ActivitiesController",
			"tag":       "GetActivities",
			"field":     result.Error.Error(),
		}).Error(errorMessage)

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}

	response := models.ActivitieResponse{
		Message: "Success",
		Status:  "Success",
		Data:    activities,
	}

	successMessage := "Activities fetched successfully"
	controller.logger.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     "info",
		"message":   successMessage,
		"source":    "ActivitiesController",
		"tag":       "GetActivities",
		"field":     "",
	}).Info(successMessage)

	c.JSON(http.StatusOK, response)
}

func (controller *ActivitiesController) GetActivity(c *gin.Context) {
	// all user can access
	db, err := utils.Connect()
	if err != nil {
		controller.logger.WithFields(logrus.Fields{
			"timestamp": time.Now().Format(time.RFC3339),
			"level":     "error",
			"message":   "Failed to connect to the database",
			"source":    "ActivitiesController",
			"tag":       "GetActivity",
			"field":     err.Error(),
		}).Error("Failed to connect to the database")

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	var activity models.Activitie
	if result := db.First(&activity, c.Param("id")); result.Error != nil {
		controller.logger.WithFields(logrus.Fields{
			"timestamp": time.Now().Format(time.RFC3339),
			"level":     "error",
			"message":   "Activity not found",
			"source":    "ActivitiesController",
			"tag":       "GetActivity",
			"field":     result.Error.Error(),
		}).Error("Activity not found")

		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Activity not found"})
		return
	}

	response := models.ActivitieResponseSingle{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	}

	controller.logger.WithFields(logrus.Fields{
		"timestamp": time.Now().Format(time.RFC3339),
		"level":     "info",
		"message":   "Activity fetched successfully",
		"source":    "ActivitiesController",
		"tag":       "GetActivity",
		"field":     "",
	}).Info("Activity fetched successfully")

	c.JSON(http.StatusOK, response)
}

func (controller *ActivitiesController) UpdateActivity(c *gin.Context) {

	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	var Activity models.Activitie
	if result := db.First(&Activity, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Activity not found"})
		return
	}

	if err := c.ShouldBind(&Activity); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if result := db.Save(&Activity); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}

	response := models.ActivitieResponseSingle{
		Status:  "Success",
		Message: "Success",
		Data:    Activity,
	}

	c.JSON(http.StatusOK, response)
}

func (controller *ActivitiesController) DeleteActivity(c *gin.Context) {
	db, err := utils.Connect()
	localZone, _ := time.LoadLocation("Asia/Bangkok")
	localTime := time.Now().In(localZone)

	if err != nil {
		controller.logger.WithFields(logrus.Fields{
			"timestamp": localTime.Format(time.RFC3339),
			"level":     "error",
			"message":   "Failed to connect to the database",
			"source":    "ActivitiesController",
			"tag":       "DeleteActivity",
			"field":     err.Error(),
		}).Error("Failed to connect to the database")

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Get the Activity ID from the URL parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		controller.logger.WithFields(logrus.Fields{
			"timestamp": localTime.Format(time.RFC3339),
			"level":     "error",
			"message":   "Invalid Activity ID",
			"source":    "ActivitiesController",
			"tag":       "DeleteActivity",
			"field":     err.Error(),
		}).Error("Invalid Activity ID")

		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid Activity ID"})
		return
	}

	// Find the Activity with the specified ID
	var activity models.Activitie
	result := db.First(&activity, id)
	if result.Error != nil {
		controller.logger.WithFields(logrus.Fields{
			"timestamp": localTime.Format(time.RFC3339),
			"level":     "error",
			"message":   "Activity not found",
			"source":    "ActivitiesController",
			"tag":       "DeleteActivity",
			"field":     result.Error.Error(),
		}).Error("Activity not found")

		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Activity not found"})
		return
	}

	// Delete the Activity
	result = db.Delete(&activity)
	if result.Error != nil {
		controller.logger.WithFields(logrus.Fields{
			"timestamp": localTime.Format(time.RFC3339),
			"level":     "error",
			"message":   "Failed to delete the activity",
			"source":    "ActivitiesController",
			"tag":       "DeleteActivity",
			"field":     result.Error.Error(),
		}).Error("Failed to delete the activity")

		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}

	controller.logger.WithFields(logrus.Fields{
		"timestamp":  localTime.Format(time.RFC3339),
		"level":      "info",
		"message":    "Activity deleted successfully",
		"source":     "ActivitiesController",
		"tag":        "DeleteActivity",
		"activityID": id,
		"userID":     activity.ID,
	}).Info("Activity deleted successfully")

	c.JSON(http.StatusOK, models.MessageResponse{Message: "Activity deleted successfully"})
}
