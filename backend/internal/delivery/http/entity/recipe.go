package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type Recipe struct {
	Uid         uuid.UUID    `json:"uid"`
	Name        string       `json:"name"`
	Products    []Product    `json:"products"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	Steps       []RecipeStep `json:"steps"`
	ImgPath     string       `json:"imgPath"`
	CookingTime int64        `json:"cookingTime"`
}

type RecipeStep struct {
	RecipeUid   uuid.UUID `json:"recipeUid"`
	Step        int64     `json:"step"`
	Description string    `json:"description"`
	ImgPath     string    `json:"imgPath"`
}

func RecipeStepFromEntity(re entity.RecipeStep) RecipeStep {
	return RecipeStep{
		RecipeUid:   re.RecipeUid,
		Step:        re.Step,
		Description: re.Description,
		ImgPath:     re.ImgPath,
	}
}

func RecipeFromEntity(re entity.Recipe) Recipe {
	products := make([]Product, len(re.Products))
	steps := make([]RecipeStep, len(re.Steps))

	for ind, product := range re.Products {
		products[ind] = ProductFromEntity(product)
	}

	for ind, step := range re.Steps {
		steps[ind] = RecipeStepFromEntity(step)
	}

	return Recipe{
		Uid:         re.Uid,
		Name:        re.Name,
		Products:    products,
		CreatedAt:   re.CreatedAt,
		UpdatedAt:   re.UpdatedAt,
		Steps:       steps,
		ImgPath:     re.ImgPath,
		CookingTime: re.CookingTime,
	}
}
