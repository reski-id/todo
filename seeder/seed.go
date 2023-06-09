package seeder

import (
	"log"
	"time"

	"todoapp/models"
	"todoapp/utils"
)

func SeedActivities() {
	db, err := utils.Connect()
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}

	// check if any activity already exists in the database
	var activity models.Activitie
	if db.First(&activity).Error == nil {
		log.Println("activities already seeded")
		return
	}

	// migrate the activitie table
	db.AutoMigrate(&models.Activitie{})

	// create some activities
	activities := []models.Activitie{
		{
			Title:     "Finish homework",
			Email:     "johndoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Buy groceries",
			Email:     "janedoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Go for a run",
			Email:     "johndoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Call mom",
			Email:     "janedoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Title:     "Write blog post",
			Email:     "johndoe@example.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	for i := range activities {
		err = db.Create(&activities[i]).Error
		if err != nil {
			log.Fatalf("failed to seed activities: %s", err.Error())
		}
	}

	log.Println("activities seeded")
}

func SeedTodos() {
	db, err := utils.Connect()
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}

	// check if any todos already exist in the database
	var todo models.Todo
	if db.First(&todo).Error == nil {
		log.Println("todos already seeded")
		return
	}

	// migrate the todo table
	db.AutoMigrate(&models.Todo{})

	// create some todos
	todos := []models.Todo{
		{
			ActivityGroupID: 1,
			Title:           "Buy groceries",
			Priority:        "high",
			IsActive:        true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ActivityGroupID: 1,
			Title:           "Do laundry",
			Priority:        "medium",
			IsActive:        true,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			ActivityGroupID: 1,
			Title:           "Write blog post",
			Priority:        "low",
			IsActive:        false,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	for i := range todos {
		err = db.Create(&todos[i]).Error
		if err != nil {
			log.Fatalf("failed to seed todos: %s", err.Error())
		}
	}

	log.Println("todos seeded")
}
