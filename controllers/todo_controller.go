package controllers

import (
	"net/http"
	"strconv"
	"todoapp/models"
	"todoapp/utils"

	"github.com/gin-gonic/gin"
)

type TodosController struct{}

func (controller TodosController) CreateTodo(c *gin.Context) {
	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	var todo models.Todo
	if err := c.ShouldBind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if result := db.Create(&todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}

	response := models.TodoSingleResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	}

	c.JSON(http.StatusCreated, response)
}

func (controller TodosController) GetTodos(c *gin.Context) {
	// all user can access
	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	var Todos []models.Todo
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset := (page - 1) * limit
	if result := db.Offset(offset).Limit(limit).Find(&Todos); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}
	var total int64
	if result := db.Model(&models.Todo{}).Count(&total); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}
	response := models.TodoResponse{
		Message: "Success",
		Status:  "Success",
		Data:    Todos,
	}
	c.JSON(http.StatusOK, response)
}

func (controller TodosController) GetTodo(c *gin.Context) {
	// all user can access
	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	var todo models.Todo
	if result := db.First(&todo, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Todo not found"})
		return
	}
	response := models.TodoSingleResponse{
		Status:  "Success",
		Message: "Success",
		Data:    todo,
	}
	c.JSON(http.StatusOK, response)
}

func (controller TodosController) UpdateTodo(c *gin.Context) {

	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	var Todo models.Todo
	if result := db.First(&Todo, c.Param("id")); result.Error != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Todo not found"})
		return
	}

	if err := c.ShouldBind(&Todo); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	if result := db.Save(&Todo); result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}

	response := models.TodoSingleResponse{
		Status:  "Success",
		Message: "Success",
		Data:    Todo,
	}

	c.JSON(http.StatusOK, response)
}

func (controller TodosController) DeleteTodo(c *gin.Context) {
	db, err := utils.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	// Get the Todo ID from the URL parameters
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid Todo ID"})
		return
	}

	// Find the Todo with the specified ID
	var Todo models.Todo
	result := db.First(&Todo, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "Todo not found"})
		return
	}

	// Delete the Todo
	result = db.Delete(&Todo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, models.MessageResponse{Message: "Todo deleted successfully"})
}
