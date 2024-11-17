package entity

import (
	uuid "github.com/satori/go.uuid"
)

type City struct {
	Uid  uuid.UUID
	Name string
}

type Address struct {
	Uid         uuid.UUID
	CityUid     uuid.UUID
	Street      string
	HouseNumber string
	Latitude    float64
	Longitude   float64
}
