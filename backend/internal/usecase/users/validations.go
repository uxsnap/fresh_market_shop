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
	if uuid.Equal(address.AddressUid, uuid.UUID{}) {
		return fmt.Errorf("address uid is empty")
	}
	if address.Apartment == 0 {
		return fmt.Errorf("apartment is empty")
	}
	if address.Floor == 0 {
		return fmt.Errorf("floor is empty")
	}

	return nil
}
