package repositoryUsers

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
	"github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/pgEntity"
)

func (r *UsersRepository) CreateDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) error {
	log.Printf("usersRepository.CreateDeliveryAddress: uid %s", address.Uid)

	if err := r.Create(ctx, pgEntity.NewDeliveryAddressRow().FromEntity(address)); err != nil {
		log.Printf("failed to create delivery address %s: %v", address.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *UsersRepository) UpdateDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) error {
	log.Printf("usersRepository.UpdateDeliveryAddress: uid %s", address.Uid)

	row := pgEntity.NewDeliveryAddressRow().FromEntity(address)
	if err := r.Update(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to update delivery address %s: %v", address.Uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *UsersRepository) DeleteDeliveryAddressByUid(ctx context.Context, uid uuid.UUID) error {
	log.Printf("usersRepository.DeleteDeliveryAddressByUid: uid %s", uid)

	row := pgEntity.NewDeliveryAddressRow().FromEntity(entity.DeliveryAddress{Uid: uid})
	if err := r.Delete(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to delete delivery address %s: %v", uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *UsersRepository) DeleteDeliveryAddressesByUserUid(ctx context.Context, userUid uuid.UUID) error {
	log.Printf("usersRepository.DeleteDeliveryAddressByUid: uid %s", userUid)

	row := pgEntity.NewDeliveryAddressRow().FromEntity(entity.DeliveryAddress{UserUid: userUid})
	if err := r.Delete(ctx, row, row.ConditionUserUidEqual()); err != nil {
		log.Printf("failed to delete delivery addresses by user uid %s: %v", userUid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (r *UsersRepository) GetDeliveryAddressByUid(ctx context.Context, uid uuid.UUID) (entity.DeliveryAddress, bool, error) {
	log.Printf("usersRepository.GetDeliveryAddressByUid: uid %s", uid)

	deliveryAddressRow := pgEntity.NewDeliveryAddressRow().FromEntity(entity.DeliveryAddress{Uid: uid})
	addressRow := pgEntity.NewAddressRow()
	cityRow := pgEntity.NewCityRow()

	selectFields := append(
		append(
			withPrefix("da", deliveryAddressRow.Columns()),
			withPrefix("ad", addressRow.Columns())...,
		), "c.name")

	selectFieldsPart := strings.Join(selectFields, ",")

	fromPart := fmt.Sprintf(
		"%s as da INNER JOIN %s as ad ON da.address_uid=ad.uid INNER JOIN %s as c ON ad.city_uid=c.uid",
		deliveryAddressRow.Table(), addressRow.Table(), cityRow.Table(),
	)

	stmt := fmt.Sprintf(`SELECT %s FROM %s WHERE da.uid=$1`, selectFieldsPart, fromPart)

	row := r.DB().QueryRow(ctx, stmt, deliveryAddressRow.Uid)

	valuesForScan := append(
		deliveryAddressRow.ValuesForScan(),
		addressRow.ValuesForScan()...,
	)
	valuesForScan = append(valuesForScan, &cityRow.Name)

	if err := row.Scan(valuesForScan...); err != nil {
		log.Printf("failed to get delivery address %s: %v", uid, err)
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.DeliveryAddress{}, false, nil
		}
		return entity.DeliveryAddress{}, false, errors.WithStack(err)
	}

	deliveryAddress := deliveryAddressRow.ToEntity()
	deliveryAddress.CityName = cityRow.Name
	deliveryAddress.StreetName = addressRow.Street
	deliveryAddress.HouseNumber = addressRow.HouseNumber
	deliveryAddress.Latitude = addressRow.Latitude
	deliveryAddress.Longitude = addressRow.Longitude

	return deliveryAddress, true, nil
}

func (r *UsersRepository) GetDeliveryAddressesByUserUid(ctx context.Context, userUid uuid.UUID) ([]entity.DeliveryAddress, error) {
	log.Printf("usersRepository.GetDeliveryAddressesByUserUid: uid %s", userUid)

	deliveryAddressRow := pgEntity.NewDeliveryAddressRow().FromEntity(entity.DeliveryAddress{UserUid: userUid})
	addressRow := pgEntity.NewAddressRow()
	cityRow := pgEntity.NewCityRow()

	selectFields := append(
		append(
			withPrefix("da", deliveryAddressRow.Columns()),
			withPrefix("ad", addressRow.Columns())...,
		), "c.name")

	selectFieldsPart := strings.Join(selectFields, ",")

	fromPart := fmt.Sprintf(
		"%s as da INNER JOIN %s as ad ON da.address_uid=ad.uid INNER JOIN %s as c ON ad.city_uid=c.uid",
		deliveryAddressRow.Table(), addressRow.Table(), cityRow.Table(),
	)

	stmt := fmt.Sprintf(`SELECT %s FROM %s WHERE da.user_uid=$1`, selectFieldsPart, fromPart)

	rows, err := r.DB().Query(ctx, stmt, deliveryAddressRow.UserUid)
	if err != nil {
		log.Printf("failed to get delivery adresses by user uid %s: %v", userUid, err)
		return nil, errors.WithStack(err)
	}

	valuesForScan := append(
		deliveryAddressRow.ValuesForScan(),
		addressRow.ValuesForScan()...,
	)
	valuesForScan = append(valuesForScan, &cityRow.Name)

	adresses := make([]entity.DeliveryAddress, 0, 10)

	for rows.Next() {
		if err := rows.Scan(valuesForScan...); err != nil {
			log.Printf("failed to scan delivery address: %v", err)
			return nil, errors.WithStack(err)
		}

		deliveryAddress := deliveryAddressRow.ToEntity()
		deliveryAddress.CityName = cityRow.Name
		deliveryAddress.StreetName = addressRow.Street
		deliveryAddress.HouseNumber = addressRow.HouseNumber
		deliveryAddress.Latitude = addressRow.Latitude
		deliveryAddress.Longitude = addressRow.Longitude

		adresses = append(adresses, deliveryAddress)
	}

	return adresses, nil
}

func withPrefix(prefix string, fields []string) []string {
	res := make([]string, 0, len(fields))
	for _, f := range fields {
		res = append(res, prefix+"."+f)
	}
	return res
}
