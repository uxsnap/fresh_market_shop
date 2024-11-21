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

type Address struct {
	Uid         uuid.UUID `json:"uid"`
	CityUid     uuid.UUID `json:"city_uid"`
	Street      string    `json:"street"`
	HouseNumber string    `json:"houseNumber"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
}

func AddressFromEntity(address entity.Address) Address {
	return Address{
		Uid:         address.Uid,
		CityUid:     address.CityUid,
		Street:      address.Street,
		HouseNumber: address.HouseNumber,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	}
}

func AddressToEntity(address Address) entity.Address {
	return entity.Address{
		Uid:         address.Uid,
		CityUid:     address.CityUid,
		Street:      address.Street,
		HouseNumber: address.HouseNumber,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
	}
}
