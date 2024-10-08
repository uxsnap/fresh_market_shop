package useCaseUsers

import (
	"context"
	"log"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

func (uc *UseCaseUsers) DeleteDeliveryAddress(ctx context.Context, uid uuid.UUID) error {
	log.Printf("ucUsers.DeleteDeliveryAddress: uid %s", uid)

	if err := uc.usersRepository.DeleteDeliveryAddressByUid(ctx, uid); err != nil {
		log.Printf("failed to delete delivery address by uid %s: %v", uid, err)
		return errors.WithStack(err)
	}
	return nil
}

func (uc *UseCaseUsers) DeleteUserDeliveryAddresses(ctx context.Context, userUid uuid.UUID) error {
	log.Printf("ucUsers.DeleteUserDeliveryAddresses: user uid %s", userUid)

	if err := uc.usersRepository.DeleteDeliveryAddressesByUserUid(ctx, userUid); err != nil {
		log.Printf("failed to delete delivery addresses by user uid %s: %v", userUid, err)
		return errors.WithStack(err)
	}
	return nil
}
