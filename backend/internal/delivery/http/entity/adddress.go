package httpEntity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type City struct {
	Uid  uuid.UUID `json:"uid"`
	Name string    `json:"name"`
}

func CityFromEntity(city entity.City) City {
	return City{
		Uid:  city.Uid,
		Name: city.Name,
	}
}

func CityToEntity(city City) entity.City {
	return entity.City{
		Uid:  city.Uid,
		Name: city.Name,
	}
}
