package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Recipe struct {
	Uid         uuid.UUID
	Name        string
	Description string
	CookingTime int64
	Products    []RecipeProduct
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type RecipeProduct struct {
	Name     string
	Quantity float64
	Measure  string
}
