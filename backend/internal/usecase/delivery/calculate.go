package useCaseDelivery

import (
	"context"
	"log"
	"math"
	"time"

	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseDelivery) CalculateDelivery(
	ctx context.Context,
	orderUid uuid.UUID,
	orderPrice int64,
	toLongitude float64,
	toLatitude float64,
) (deliveryPrice int64, deliveryTime time.Duration, err error) {
	log.Printf("usecaseDelivery.CalculateDelivery: order uid %s", orderUid)

	toLongRad, toLatRad, fromLongRad, fromLatRad := degToRad(toLongitude), degToRad(toLatitude), degToRad(fromLongitude), degToRad(fromLatitude)

	dLong := fromLongRad - toLongRad
	dist := math.Atan(math.Sqrt(math.Pow(math.Cos(toLatRad)*math.Sin(dLong), 2)+math.Pow(math.Cos(fromLatRad)*math.Sin(toLatRad)-math.Sin(fromLatRad)*math.Cos(toLatRad)*math.Cos(dLong), 2)) / (math.Sin(fromLatRad)*math.Sin(toLatRad) + math.Cos(fromLatRad)*math.Cos(toLatRad)*math.Cos(dLong)))
	dist *= earthRadiusM

	t := dist / courierSpeed
	deliveryTime = time.Duration(int64(t)) * time.Minute
	deliveryTime += 5 * time.Minute // 5 минут на подмыться

	deliveryPrice = int64(dist / 100 * priceForHungredMetres)

	if deliveryPrice < minDeliveryPrice {
		deliveryPrice = minDeliveryPrice
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
