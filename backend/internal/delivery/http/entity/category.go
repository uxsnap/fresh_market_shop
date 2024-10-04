package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type Category struct {
	Uid         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func CategoryFromEntity(category entity.Category) Category {
	return Category{
		Uid:         category.Uid,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}

func CategoryToEntity(category Category) entity.Category {
	return entity.Category{
		Uid:         category.Uid,
		Name:        category.Name,
		Description: category.Description,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
	}
}
