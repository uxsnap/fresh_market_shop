package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type Recipe struct {
	Uid         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CookingTime int64     `json:"cookingTime"`
	Ccal        int64     `json:"ccal"`
}

type RecipeStep struct {
	RecipeUid   uuid.UUID `json:"recipeUid"`
	Step        int64     `json:"step"`
	Description string    `json:"description"`
}

func RecipeStepFromEntity(re entity.RecipeStep) RecipeStep {
	return RecipeStep{
		RecipeUid:   re.RecipeUid,
		Step:        re.Step,
		Description: re.Description,
	}
}

func RecipeFromEntity(re entity.Recipe) Recipe {
	return Recipe{
		Uid:         re.Uid,
		Name:        re.Name,
		CreatedAt:   re.CreatedAt,
		UpdatedAt:   re.UpdatedAt,
		CookingTime: re.CookingTime,
		Ccal:        re.Ccal,
	}
}
