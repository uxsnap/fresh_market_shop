package repositoryUsers

import (
	"context"
	"log"

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

	row := pgEntity.NewDeliveryAddressRow().FromEntity(entity.DeliveryAddress{Uid: uid})
	if err := r.GetOne(ctx, row, row.ConditionUidEqual()); err != nil {
		log.Printf("failed to get delivery address %s: %v", uid, err)
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.DeliveryAddress{}, false, nil
		}
		return entity.DeliveryAddress{}, false, errors.WithStack(err)
	}
	return row.ToEntity(), true, nil
}

func (r *UsersRepository) GetDeliveryAddressesByUserUid(ctx context.Context, userUid uuid.UUID) ([]entity.DeliveryAddress, error) {
	log.Printf("usersRepository.GetDeliveryAddressesByUserUid: uid %s", userUid)

	row := pgEntity.NewDeliveryAddressRow().FromEntity(entity.DeliveryAddress{UserUid: userUid})
	rows := pgEntity.NewDeliveryAddressRows()

	if err := r.GetSome(ctx, row, rows, row.ConditionUserUidEqual()); err != nil {
		log.Printf("failed to get user %s delivery addresses: %v", userUid, err)
		return nil, errors.WithStack(err)
	}

	return rows.ToEntity(), nil
}
