package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type DeliveryAddress struct {
	Uid         uuid.UUID `json:"uid"`
	UserUid     uuid.UUID `json:"userUid"`
	AddressUid  uuid.UUID `json:"addressUid"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CityName    string    `json:"cityName"`
	StreetName  string    `json:"streetName"`
	HouseNumber string    `json:"houseNumber"`
	Floor       int64     `json:"floor"`
	Entrance    int64     `json:"entrance"`
	Apartment   int64     `json:"apartment"`
	Code        int64     `json:"code"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func DeliveryAddressFromEntity(address entity.DeliveryAddress) DeliveryAddress {
	return DeliveryAddress{
		Uid:         address.Uid,
		UserUid:     address.UserUid,
		AddressUid:  address.AddressUid,
		Latitude:    address.Latitude,
		Longitude:   address.Longitude,
		CityName:    address.CityName,
		StreetName:  address.StreetName,
		HouseNumber: address.HouseNumber,
		Floor:       address.Floor,
		Entrance:    address.Entrance,
		Apartment:   address.Apartment,
		Code:        address.Code,
		CreatedAt:   address.CreatedAt,
		UpdatedAt:   address.UpdatedAt,
	}
}

func DeliveryAddressToEntity(address DeliveryAddress) entity.DeliveryAddress {
	return entity.DeliveryAddress{
		Uid:        address.Uid,
		UserUid:    address.UserUid,
		AddressUid: address.AddressUid,
		Apartment:  address.Apartment,
		Entrance:   address.Entrance,
		Floor:      address.Floor,
		Code:       address.Code,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}
}
