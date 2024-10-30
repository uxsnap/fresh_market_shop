package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Recipe struct {
	Uid         uuid.UUID
	Name        string
	CookingTime int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ImgPath     string
	Steps       []RecipeStep
	Products    []Product
}

type RecipeStep struct {
	RecipeUid   uuid.UUID
	Step        int64
	Description string
	ImgPath     string
}
