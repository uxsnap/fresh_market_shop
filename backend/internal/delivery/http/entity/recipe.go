package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Recipe struct {
	Uid         uuid.UUID     `json:"uid"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Products    []Product     `json:"products"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Steps       []RecipeSteps `json:"steps"`
	ImgPath     string        `json:"imgPath"`
}

type RecipeSteps struct {
	RecipeUid   uuid.UUID `json:"recipeUid"`
	Step        int64     `json:"step"`
	Description string    `json:"description"`
	ImgPath     string    `json:"imgPath"`
}
