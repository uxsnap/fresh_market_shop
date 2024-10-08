package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type DeliveryAddress struct {
	Uid         uuid.UUID `json:"uid"`
	UserUid     uuid.UUID `json:"userUid"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CityName    string    `json:"cityName"`
	StreetName  string    `json:"streetName"`
	HouseNumber int64     `json:"houseNumber"`
	Building    int64     `json:"building"`
	Floor       int64     `json:"floor"`
	Apartment   int64     `json:"apartment"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func DeliveryAddressFromEntity(address entity.DeliveryAddress) DeliveryAddress {
	return DeliveryAddress{
		Uid:         address.Uid,
		UserUid:     address.UserUid,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
		CityName:    address.CityName,
		StreetName:  address.StreetName,
		HouseNumber: address.HouseNumber,
		Building:    address.Building,
		Floor:       address.Floor,
		Apartment:   address.Apartment,
		CreatedAt:   address.CreatedAt,
		UpdatedAt:   address.UpdatedAt,
	}
}

func DeliveryAddressToEntity(address DeliveryAddress) entity.DeliveryAddress {
	return entity.DeliveryAddress{
		Uid:         address.Uid,
		UserUid:     address.UserUid,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
		CityName:    address.CityName,
		StreetName:  address.StreetName,
		HouseNumber: address.HouseNumber,
		Building:    address.Building,
		Floor:       address.Floor,
		Apartment:   address.Apartment,
		CreatedAt:   address.CreatedAt,
		UpdatedAt:   address.UpdatedAt,
	}
}
