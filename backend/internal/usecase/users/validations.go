package useCaseUsers

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func validateDeliveryAddress(address entity.DeliveryAddress) error {
	if uuid.Equal(address.UserUid, uuid.UUID{}) {
		return fmt.Errorf("user uid is empty")
	}
	if len(address.CityName) == 0 {
		return fmt.Errorf("city name is empty")
	}
	if len(address.StreetName) == 0 {
		return fmt.Errorf("street name is empty")
	}
	if address.Building == 0 {
		return fmt.Errorf("building is empty")
	}
	if address.Latitude == 0 {
		return fmt.Errorf("latitude is empty")
	}
	if address.Longitude == 0 {
		return fmt.Errorf("longitude is empty")
	}

	return nil
}
