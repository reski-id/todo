package seeder

import (
	"todoapp/models"
	"todoapp/utils"
)

func CreateMigration() {
	db, err := utils.Connect()
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Auto migrate all entities
	db.AutoMigrate(&models.Activitie{}, &models.Todo{})
}
