package useCaseDelivery

import (
	"context"
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (uc *UseCaseDelivery) CalculateDelivery(
	ctx context.Context,
	userUid uuid.UUID,
	orderUid uuid.UUID,
	deliveryAddressUid uuid.UUID,
) (deliveryPrice int64, deliveryTime time.Duration, err error) {
	log.Printf("usecaseDelivery.CalculateDelivery: order uid %s", orderUid)

	user, isFound, err := uc.usersService.GetUser(ctx, userUid)
	if err != nil {
		log.Printf("failed to get user by uid %s: %v", userUid, err)
		return 0, 0, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("user with uid %s not found", userUid)
		return 0, 0, errors.New("user not found")
	}

	order, isFound, err := uc.ordersService.GetOrder(ctx, orderUid)
	if err != nil {
		log.Printf("failed to get order by uid %s: %v", orderUid, err)
		return 0, 0, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("order with uid %s not found", userUid)
		return 0, 0, errors.New("order not found")
	}

	deliveryAddress, isFound, err := uc.usersService.GetDeliveryAddress(ctx, deliveryAddressUid)
	if err != nil {
		log.Printf("failed to get delivery address by uid %s: %v", deliveryAddressUid, err)
		return 0, 0, errors.WithStack(err)
	}
	if !isFound {
		log.Printf("delivery address with uid %s not found", deliveryAddressUid)
		return 0, 0, errors.New("delivery address not found")
	}

	toLongRad, toLatRad, fromLongRad, fromLatRad := degToRad(deliveryAddress.Longitude), degToRad(deliveryAddress.Latitude), degToRad(fromLongitude), degToRad(fromLatitude)

	dLong := fromLongRad - toLongRad
	dist := math.Atan(math.Sqrt(math.Pow(math.Cos(toLatRad)*math.Sin(dLong), 2)+math.Pow(math.Cos(fromLatRad)*math.Sin(toLatRad)-math.Sin(fromLatRad)*math.Cos(toLatRad)*math.Cos(dLong), 2)) / (math.Sin(fromLatRad)*math.Sin(toLatRad) + math.Cos(fromLatRad)*math.Cos(toLatRad)*math.Cos(dLong)))
	dist *= earthRadiusM

	t := dist / courierSpeed
	deliveryTime = time.Duration(int64(t)) * time.Minute
	deliveryTime += 5 * time.Minute // 5 минут на подмыться

	deliveryPrice = int64(dist/100*priceForHungredMetres) + int64(float64(order.Sum)*0.05)

	if deliveryPrice < minDeliveryPrice {
		deliveryPrice = minDeliveryPrice
	}

	addressParts := []string{
		deliveryAddress.CityName,
		deliveryAddress.StreetName,
		strconv.Itoa(int(deliveryAddress.HouseNumber)),
		strconv.Itoa(int(deliveryAddress.Building)),
		strconv.Itoa(int(deliveryAddress.Floor)),
		strconv.Itoa(int(deliveryAddress.Apartment)),
	}

	addressStr := strings.Join(addressParts, " ")

	delivery := entity.Delivery{
		Uid:           uuid.NewV4(),
		OrderUid:      orderUid,
		FromLongitude: fromLongitude,
		FromLatitude:  fromLatitude,
		ToLongitude:   deliveryAddress.Longitude,
		ToLatitude:    deliveryAddress.Latitude,
		Address:       addressStr,
		Receiver:      user.FirstName + " " + user.LastName,
		Time:          int64(deliveryTime),
		Price:         deliveryPrice,
		CreatedAt:     time.Now().UTC(),
	}

	if err := uc.deliveryRepository.CreateDelivery(ctx, delivery); err != nil {
		log.Printf("failed to create delivery: %v", err)
		return 0, 0, errors.WithStack(err)
	}

	return deliveryPrice, deliveryTime, nil
}

const (
	earthRadiusM = 6372795.0
	courierSpeed = 500.0 // метры в минуту
	// это переедет в тарифы:
	priceForHungredMetres = 50.0
	minDeliveryPrice      = 150
)

func degToRad(deg float64) float64 {
	return deg * math.Pi / 180
}
