package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Uid       uuid.UUID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DeliveryAddress struct {
	Uid         uuid.UUID
	UserUid     uuid.UUID
	Latitude    float64
	Longitude   float64
	CityName    string
	StreetName  string
	HouseNumber int64
	Building    int64
	Floor       int64
	Apartment   int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
