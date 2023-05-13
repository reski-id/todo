package controllers

import (
	"net/http"
	"strconv"
	"todoapp/models"
	"todoapp/utils"

	"github.com/gin-gonic/gin"
)

type ActivitiesController struct{}

func (controller ActivitiesController) CreateActivity(c *gin.Context) {
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

func (controller ActivitiesController) GetActivities(c *gin.Context) {
	// all user can access
	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	var Activitys []models.Activitie
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset := (page - 1) * limit
	if result := db.Offset(offset).Limit(limit).Find(&Activitys); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}
	var total int64
	if result := db.Model(&models.Activitie{}).Count(&total); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}
	response := models.ActivitieResponse{
		Message: "Success",
		Status:  "Success",
		Data:    Activitys,
	}
	c.JSON(http.StatusOK, response)
}

func (controller ActivitiesController) GetActivity(c *gin.Context) {
	// all user can access
	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	var activity models.Activitie
	if result := db.First(&activity, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Activity not found"})
		return
	}
	response := models.ActivitieResponseSingle{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	}
	c.JSON(http.StatusOK, response)
}

func (controller ActivitiesController) UpdateActivity(c *gin.Context) {

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

func (controller ActivitiesController) DeleteActivity(c *gin.Context) {
	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Get the Activity ID from the URL parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid Activity ID"})
		return
	}

	// Find the Activity with the specified ID
	var Activity models.Activitie
	result := db.First(&Activity, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Activity not found"})
		return
	}

	// Delete the Activity
	result = db.Delete(&Activity)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, models.MessageResponse{Message: "Activity deleted successfully"})
}
