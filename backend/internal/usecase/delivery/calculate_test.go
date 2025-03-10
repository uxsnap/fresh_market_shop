package useCaseDelivery

import (
	"fmt"
	"testing"
	"time"
)

type CalcDistCase struct {
	FromLatitude  float64
	FromLongitude float64

	ToLongitude float64
	ToLatitude  float64
}

/*
	59.932472, 30.347427 Невский проспект 47 (1,25 км +-)
	59.956090, 30.298274 улица Лизы Чайкиной (3,9км +- )

*/

func TestCalcDist(t *testing.T) {

	cases := []CalcDistCase{
		{
			FromLatitude:  fromLatitude,
			FromLongitude: fromLongitude,
			ToLongitude:   59.932472,
			ToLatitude:    30.347427,
		},
		{
			FromLatitude:  fromLatitude,
			FromLongitude: fromLongitude,
			ToLongitude:   59.956090,
			ToLatitude:    30.298274,
		},
	}

	for _, c := range cases {
		dist := calcDist(c.ToLongitude, c.ToLatitude, c.FromLongitude, c.FromLatitude)
		fmt.Println("dist: ", dist)

		t := dist / courierSpeed

		deliveryTime := time.Duration(int64(t)) * time.Minute
		deliveryTime *= 2 //  с запасом

		fmt.Println("delivery time: ", deliveryTime)

		// cумма заказа 1000 рублей
		deliveryPrice := int64(dist/100*priceForHundredMetres) + int64(float64(1000)*0.05)

		fmt.Println("delivery price dirty: ", deliveryPrice)
		if deliveryPrice < minDeliveryPrice {
			deliveryPrice = minDeliveryPrice
		} else if deliveryPrice > maxDeliveryPrice {
			deliveryPrice = maxDeliveryPrice
		}

		fmt.Println("delivery price: ", deliveryPrice)
	}
}
