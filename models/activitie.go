package models

import (
	"time"
)

type Activitie struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" form:"title"`
	Email     string    `json:"email" form:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ActivitieResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    []Activitie `json:"data"`
}

type ActivitieResponseSingle struct {
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    Activitie `json:"data"`
}
